package exasol_test_setup_abstraction_go

import (
	"bytes"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/exasol/exasol-driver-go"
)

type TestSetupAbstraction struct {
	server         *exec.Cmd
	serverEndpoint string
	stopped        *bool
	stoppedMutex   *sync.Mutex
	errorStream    *bytes.Buffer
}

const serverVersion = "0.2.0"
const serverJar = "exasol-test-setup-abstraction-server-" + serverVersion + ".jar"

func Create(configFilePath string) (*TestSetupAbstraction, error) {
	serverPath, err := downloadServerIfNotPresent()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Starting server jar %q with configuration %q...\n", serverPath, configFilePath)
	serverProcess := exec.Command("java", "-jar", serverPath, configFilePath)
	var output, errorStream bytes.Buffer
	serverProcess.Stdout = &output
	serverProcess.Stderr = &errorStream
	stoppedMutex := &sync.Mutex{}
	stopped := false
	go waitForServer(serverProcess, &errorStream, &stopped, stoppedMutex)
	port, err := getServerPort(&stopped, &output, &errorStream)
	if err != nil {
		return nil, err
	}
	serverEndpoint := fmt.Sprintf("http://localhost:%v/", port)
	return &TestSetupAbstraction{server: serverProcess, serverEndpoint: serverEndpoint, stopped: &stopped, stoppedMutex: stoppedMutex, errorStream: &errorStream}, nil
}

func downloadServerIfNotPresent() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir. Cause: %v", err.Error())
	}
	serverDir := path.Join(homeDir, ".test-setup-abstraction-server")
	if _, err := os.Stat(serverDir); os.IsNotExist(err) {
		err := os.Mkdir(serverDir, 0700)
		if err != nil {
			return "", fmt.Errorf("failed to create directory for the test-setup-abstraction-server. Cause: %v", err.Error())
		}
	}
	serverFile := path.Join(serverDir, serverJar)
	if _, err := os.Stat(serverFile); os.IsNotExist(err) {
		out, err := os.Create(serverFile)
		defer func() {
			err = out.Close()
			if err != nil {
				panic(fmt.Sprintf("failed to close server file. Cause: %v", err.Error()))
			}
		}()
		if err != nil {
			return "", fmt.Errorf("failed to create file for downloading test-setup-abstraction-server. Cause: %v", err.Error())
		}
		resp, err := http.Get("https://github.com/exasol/exasol-test-setup-abstraction-server/releases/download/" + serverVersion + "/" + serverJar)
		if err != nil {
			return "", fmt.Errorf("failed to download test-setup-abstraction-server. Cause: %v", err.Error())
		}
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return "", fmt.Errorf("failed to download test-setup-abstraction-server. Cause: %v", err.Error())
		}
	}
	return serverFile, nil
}

func getServerPort(stopped *bool, output *bytes.Buffer, errorStream *bytes.Buffer) (int, error) {
	for counter := 0; counter < 10; counter++ {
		pattern := regexp.MustCompile("Server running on port: (\\d+)\n")
		result := pattern.FindSubmatch(output.Bytes())
		if len(result) != 0 {
			portString := string(result[1])
			port, err := strconv.ParseInt(portString, 10, 32)
			if err != nil {
				return -1, fmt.Errorf("failed to parse port %q", portString)
			}
			return int(port), nil
		}
		if !*stopped {
			time.Sleep(1 * time.Second)
		}
	}
	return -1, fmt.Errorf("failed to start server. The server did not print a port number. Output: %q, error stream: %q", output, errorStream)
}

func waitForServer(serverProcess *exec.Cmd, errorStream *bytes.Buffer, stopped *bool, stoppedMutex *sync.Mutex) {
	err := serverProcess.Run()
	if err != nil && !isStopped(stopped, stoppedMutex) { // after we killed the thread we expect an error
		fmt.Println(errorStream.String())
	}
	stoppedMutex.Lock()
	*stopped = true
	stoppedMutex.Unlock()
}

func isStopped(stopped *bool, stoppedMutex *sync.Mutex) bool {
	stoppedMutex.Lock()
	defer stoppedMutex.Unlock()
	return *stopped
}

func (testSetup *TestSetupAbstraction) Stop() error {
	testSetup.stoppedMutex.Lock()
	*testSetup.stopped = true
	testSetup.stoppedMutex.Unlock()
	err := testSetup.server.Process.Signal(os.Kill)
	if err != nil {
		return fmt.Errorf("failed to stop test-setup-abstraction server. Cause: %v", err.Error())
	}
	return nil
}

func (testSetup *TestSetupAbstraction) getConnectionInfo() (*ConnectionInfo, error) {
	var connectionDetails ConnectionInfo
	err := testSetup.makeApiRequest("GET", "connectionInfo", &connectionDetails, nil)
	if err != nil {
		return nil, err
	}
	return &connectionDetails, nil
}

func (testSetup *TestSetupAbstraction) makeApiRequest(method string, path string, jsonResult interface{}, payload url.Values) error {
	client := http.Client{}
	req, err := http.NewRequest(method, testSetup.serverEndpoint+path, strings.NewReader(payload.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create http request for the server. Cause: %v", err.Error())
	}
	if payload != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute %v %v from test-setup-abstraction server. Cause %v", method, path, err.Error())
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body. Cause %v", err.Error())
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("request failed with status %d (%v). Response: %q", response.StatusCode, response.Status, body)
	}
	err = json.Unmarshal(body, &jsonResult)
	if err != nil {
		return fmt.Errorf("invalid JSON response:\n%v\nCause: %v", string(body), err.Error())
	}
	return nil
}

type ConnectionInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (testSetup *TestSetupAbstraction) CreateConnection() (*sql.DB, error) {
	return testSetup.CreateConnectionWithConfig(true)
}

func (testSetup *TestSetupAbstraction) CreateConnectionWithConfig(autocommit bool) (*sql.DB, error) {
	connectionDetails, err := testSetup.getConnectionInfo()
	if err != nil {
		return nil, err
	}
	connection, err := sql.Open("exasol", exasol.NewConfig(connectionDetails.User, connectionDetails.Password).
		Port(connectionDetails.Port).
		Host(connectionDetails.Host).
		ValidateServerCertificate(false).
		Autocommit(autocommit).
		String())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database. Cause %v", err.Error())
	}
	return connection, nil
}

func (testSetup *TestSetupAbstraction) MakeDatabaseTcpServiceAccessibleFromLocalhost(databasePort int) ([]int, error) {
	var ports []int
	err := testSetup.makeApiRequest("POST", "makeDatabaseTcpServiceAccessibleFromLocalhost", &ports, url.Values{
		"databasePort": {strconv.Itoa(databasePort)},
	})
	if err != nil {
		return nil, err
	}
	return ports, nil
}

func (testSetup *TestSetupAbstraction) MakeLocalTcpServiceAccessibleFromDatabase(localPort int) (*ServiceAddress, error) {
	var serviceAddress ServiceAddress
	err := testSetup.makeApiRequest("POST", "makeLocalTcpServiceAccessibleFromDatabase", &serviceAddress, url.Values{
		"localPort": {strconv.Itoa(localPort)},
	})
	if err != nil {
		return nil, err
	}
	return &serviceAddress, nil
}

func (testSetup *TestSetupAbstraction) MakeTcpServiceAccessibleFromDatabase(serviceAddress ServiceAddress) (*ServiceAddress, error) {
	var localAddress ServiceAddress
	err := testSetup.makeApiRequest("POST", "makeTcpServiceAccessibleFromDatabase", &localAddress, url.Values{
		"hostName": {string(serviceAddress.HostName)},
		"port":     {strconv.Itoa(serviceAddress.Port)},
	})
	if err != nil {
		return nil, err
	}
	return &localAddress, err
}

func (testSetup TestSetupAbstraction) UploadFile(localPath string, remoteName string) error {
	return testSetup.makeApiRequest("POST", "bfs/uploadFile", &successResult{}, url.Values{
		"localPath":  {localPath},
		"remoteName": {remoteName},
	})
}

func (testSetup TestSetupAbstraction) UploadStringContent(stringContent string, remoteName string) error {
	return testSetup.makeApiRequest("POST", "bfs/uploadStringContent", &successResult{}, url.Values{
		"stringContent": {stringContent},
		"remoteName":    {remoteName},
	})
}

type downloadFileAsStringResult struct {
	Content string `json:"content"`
}

func (testSetup TestSetupAbstraction) DownloadFileAsString(path string) (string, error) {
	result := downloadFileAsStringResult{}
	err := testSetup.makeApiRequest("GET", "bfs/downloadFileAsString?path="+url.QueryEscape(path), &result, url.Values{})
	if err != nil {
		return "", err
	}
	return result.Content, nil
}

func (testSetup TestSetupAbstraction) DownloadFile(remotePath string, localPath string) error {
	return testSetup.makeApiRequest("GET", "bfs/downloadFile?remotePath="+url.QueryEscape(remotePath)+"&localPath="+url.QueryEscape(localPath), &successResult{}, url.Values{})
}

func (testSetup TestSetupAbstraction) DeleteFile(path string) error {
	return testSetup.makeApiRequest("DELETE", "bfs/deleteFile", &successResult{}, url.Values{
		"path": {path},
	})
}

func (testSetup TestSetupAbstraction) ListFiles(path string) ([]string, error) {
	result := &listResult{}
	err := testSetup.makeApiRequest("GET", "bfs/listFiles?path="+url.QueryEscape(path), result, url.Values{
		"path": {path},
	})
	return result.Files, err
}

type listResult struct {
	Files []string `json:"files"`
}

type successResult struct {
	Success bool `json:"success"`
}

type ServiceAddress struct {
	HostName string `json:"hostName"`
	Port     int    `json:"port"`
}
