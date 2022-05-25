package exasol_test_setup_abstraction_go

import (
	"bytes"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/exasol/exasol-driver-go"
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
)

type TestSetupAbstraction struct {
	server         *exec.Cmd
	serverEndpoint string
	stopped        *bool
	stoppedMutex   *sync.Mutex
}

const serverVersion = "0.1.1"
const serverJar = "exasol-test-setup-abstraction-server-" + serverVersion + ".jar"

func Create(configFilePath string) TestSetupAbstraction {
	serverPath := downloadServerIfNotPresent()
	serverProcess := exec.Command("java", "-jar", serverPath, configFilePath)
	var output, errorStream bytes.Buffer
	serverProcess.Stdout = &output
	serverProcess.Stderr = &errorStream
	stoppedMutex := &sync.Mutex{}
	stopped := false
	go waitForServer(serverProcess, &errorStream, &stopped, stoppedMutex)
	port := getServerPort(&output, &errorStream)
	serverEndpoint := fmt.Sprintf("http://localhost:%v/", port)
	return TestSetupAbstraction{server: serverProcess, serverEndpoint: serverEndpoint, stopped: &stopped, stoppedMutex: stoppedMutex}
}

func downloadServerIfNotPresent() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("failed to get home dir. Cause: %v", err.Error()))
	}
	serverDir := path.Join(homeDir, ".test-setup-abstraction-server")
	if _, err := os.Stat(serverDir); os.IsNotExist(err) {
		err := os.Mkdir(serverDir, 0700)
		if err != nil {
			panic(fmt.Sprintf("failed to create directory for the test-setup-abstraction-server. Cause: %v", err.Error()))
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
			panic(fmt.Sprintf("failed to create file for downloading test-setup-abstraction-server. Cause: %v", err.Error()))
		}
		resp, err := http.Get("https://github.com/exasol/exasol-test-setup-abstraction-server/releases/download/" + serverVersion + "/" + serverJar)
		if err != nil {
			panic(fmt.Sprintf("failed to download test-setup-abstraction-server. Cause: %v", err.Error()))
		}
		defer resp.Body.Close()
		_, err = io.Copy(out, resp.Body)
		if err != nil {
			panic(fmt.Sprintf("failed to download test-setup-abstraction-server. Cause: %v", err.Error()))
		}
	}
	return serverFile
}

func getServerPort(output *bytes.Buffer, errorStream *bytes.Buffer) int {
	for counter := 0; counter < 500; counter++ {
		pattern := regexp.MustCompile("Server running on port: (\\d+)\n")
		result := pattern.FindSubmatch(output.Bytes())

		if len(result) != 0 {
			portString := string(result[1])
			port, err := strconv.ParseInt(portString, 10, 32)
			if err != nil {
				panic(err)
			}
			return int(port)
		}
		time.Sleep(1 * time.Second)
	}
	fmt.Println(errorStream.String())
	fmt.Println(output.String())
	panic("failed to start server. The server did not print a port number.")
}

func waitForServer(serverProcess *exec.Cmd, errorStream *bytes.Buffer, stopped *bool,
	stoppedMutex *sync.Mutex) {
	err := serverProcess.Run()
	if err != nil && !isStopped(stopped, stoppedMutex) { //after we killed the thread we expect an error
		fmt.Println(errorStream.String())
		panic(fmt.Sprintf("failed to start test-setup-abstraction server. Cause: %v", err.Error()))
	}
}

func isStopped(stopped *bool, stoppedMutex *sync.Mutex) bool {
	stoppedMutex.Lock()
	defer stoppedMutex.Unlock()
	return *stopped
}

func (testSetup *TestSetupAbstraction) Stop() {
	testSetup.stoppedMutex.Lock()
	*testSetup.stopped = true
	testSetup.stoppedMutex.Unlock()
	err := testSetup.server.Process.Signal(os.Kill)
	if err != nil {
		panic(fmt.Sprintf("failed to stop test-setup-abstraction server. Cause: %v", err.Error()))
	}
}

func (testSetup *TestSetupAbstraction) getConnectionInfo() *ConnectionInfo {
	var connectionDetails ConnectionInfo
	testSetup.makeApiRequest("GET", "connectionInfo", &connectionDetails, nil)
	return &connectionDetails
}

func (testSetup *TestSetupAbstraction) makeApiRequest(method string, path string, jsonResult interface{}, payload url.Values) {
	client := http.Client{}
	req, err := http.NewRequest(method, testSetup.serverEndpoint+path, strings.NewReader(payload.Encode()))
	if err != nil {
		panic(fmt.Sprintf("failed to create http request for the server. Cause: %v", err.Error()))
	}
	if payload != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	response, err := client.Do(req)
	if err != nil {
		panic(fmt.Sprintf("failed to execute %v %v from test-setup-abstraction server. Cause %v", method, path, err.Error()))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(body, &jsonResult)
	if err != nil {
		panic(fmt.Sprintf("invalid JSON response:\n%v\nCause: %v", string(body), err.Error()))
	}
}

type ConnectionInfo struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func (testSetup *TestSetupAbstraction) CreateConnection() *sql.DB {
	connectionDetails := testSetup.getConnectionInfo()
	connection, err := sql.Open("exasol", exasol.NewConfig(connectionDetails.User, connectionDetails.Password).
		Port(connectionDetails.Port).
		Host(connectionDetails.Host).
		ValidateServerCertificate(false).
		String())
	if err != nil {
		panic(fmt.Sprintf("failed to connect to the database. Cause %v", err.Error()))
	}
	return connection
}

func (testSetup *TestSetupAbstraction) MakeDatabaseTcpServiceAccessibleFromLocalhost(databasePort int) []int {
	var ports []int
	testSetup.makeApiRequest("POST", "makeDatabaseTcpServiceAccessibleFromLocalhost", &ports, url.Values{
		"databasePort": {strconv.Itoa(databasePort)},
	})
	return ports
}

func (testSetup *TestSetupAbstraction) MakeLocalTcpServiceAccessibleFromDatabase(localPort int) ServiceAddress {
	var serviceAddress ServiceAddress
	testSetup.makeApiRequest("POST", "makeLocalTcpServiceAccessibleFromDatabase", &serviceAddress, url.Values{
		"localPort": {strconv.Itoa(localPort)},
	})
	return serviceAddress
}

func (testSetup *TestSetupAbstraction) MakeTcpServiceAccessibleFromDatabase(serviceAddress ServiceAddress) ServiceAddress {
	var localAddress ServiceAddress
	testSetup.makeApiRequest("POST", "makeTcpServiceAccessibleFromDatabase", &localAddress, url.Values{
		"hostName": {string(serviceAddress.HostName)},
		"port":     {strconv.Itoa(serviceAddress.Port)},
	})
	return localAddress
}

type ServiceAddress struct {
	HostName string `json:"hostName"`
	Port     int    `json:"port"`
}
