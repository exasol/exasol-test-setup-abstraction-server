package exasol_test_setup_abstraction_go

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/exasol/exasol-driver-go"
)

type TestSetupAbstraction struct {
	server *serverProcess
}

const serverVersion = "0.3.8"

// Create creates a new Exasol test setup with the given path to the config file
// and starts a local server.
// If the file does not exists, a local Docker container will be started.
func startTestSetupAbstraction(config Builder) (*TestSetupAbstraction, error) {
	server, err := startServer(serverVersion, config)
	if err != nil {
		return nil, err
	}
	return &TestSetupAbstraction{server: server}, nil
}

// Stop stops the local server. Docker containers will keep running if reuse is enabled in ~/.testcontainers.properties.
func (testSetup *TestSetupAbstraction) Stop() error {
	return testSetup.server.stop()
}

// GetConnectionInfo returns details required to connect to the Exasol database.
func (testSetup *TestSetupAbstraction) GetConnectionInfo() (*ConnectionInfo, error) {
	var connectionDetails ConnectionInfo
	err := testSetup.makeApiRequest("GET", "connectionInfo", &connectionDetails, nil)
	if err != nil {
		return nil, err
	}
	return &connectionDetails, nil
}

func (testSetup *TestSetupAbstraction) makeApiRequest(method string, path string, jsonResult interface{}, payload url.Values) error {
	//nolint:exhaustruct // default values are OK
	client := http.Client{}
	req, err := http.NewRequestWithContext(context.Background(), method, testSetup.server.serverEndpoint+path, strings.NewReader(payload.Encode()))
	if err != nil {
		return fmt.Errorf("failed to create http request for the server. Cause: %w", err)
	}
	if payload != nil {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	response, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to execute %v %v from test-setup-abstraction server. Cause %w", method, path, err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body. Cause %w", err)
	}
	if response.StatusCode != 200 {
		return fmt.Errorf("request failed with status %d (%v). Response: %q", response.StatusCode, response.Status, body)
	}
	err = json.Unmarshal(body, &jsonResult)
	if err != nil {
		return fmt.Errorf("invalid JSON response:\n%v\nCause: %w", string(body), err)
	}
	return nil
}

// Contains information required for connecting to an exasol database.
type ConnectionInfo struct {
	Host     string `json:"host"`     // Host name
	Port     int    `json:"port"`     // Port number
	User     string `json:"user"`     // User name
	Password string `json:"password"` // Password
}

// CreateConnection returns a new database connection with autocommit enabled.
func (testSetup *TestSetupAbstraction) CreateConnection() (*sql.DB, error) {
	return testSetup.CreateConnectionWithConfig(true)
}

// CreateConnection returns a new database connection with autocommit on or off.
func (testSetup *TestSetupAbstraction) CreateConnectionWithConfig(autocommit bool) (*sql.DB, error) {
	connectionDetails, err := testSetup.GetConnectionInfo()
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
		return nil, fmt.Errorf("failed to connect to the database. Cause %w", err)
	}
	return connection, nil
}

// MakeDatabaseTcpServiceAccessibleFromLocalhost makes the port of a database node available from localhost.
// If the target is a cluster, this methods sets up a redirect for each node and returns multiple local port numbers.
// You can use this method for example to connect from your test PC to a profiling agent that listens on a TCP socket in an UDF.
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

// MakeLocalTcpServiceAccessibleFromDatabase makes a local TCP service available from within the Exasol database.
// You can use this method for example for accessing a local s3 bucket implementation from inside the Exasol database.
// Another example is to connect from UDFs to a debugger running on the test PC.
// Returns the address under which the service is available from within the exasol database (same port).
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

// MakeTcpServiceAccessibleFromDatabase makes a given TCP service available inside of the Exasol database and returns the modified service address.
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

// UploadFile uploads a local file to the default BucketFS bucket.
func (testSetup TestSetupAbstraction) UploadFile(localPath string, remoteName string) error {
	return testSetup.makeApiRequest("POST", "bfs/uploadFile", &successResult{Success: true}, url.Values{
		"localPath":  {localPath},
		"remoteName": {remoteName},
	})
}

// UploadStringContent uploads the given string content to a file in the default BucketFS bucket.
func (testSetup TestSetupAbstraction) UploadStringContent(stringContent string, remoteName string) error {
	return testSetup.makeApiRequest("POST", "bfs/uploadStringContent", &successResult{Success: true}, url.Values{
		"stringContent": {stringContent},
		"remoteName":    {remoteName},
	})
}

type downloadFileAsStringResult struct {
	Content string `json:"content"`
}

// DownloadFileAsString downloads a file from the default BucketFS bucket.
func (testSetup TestSetupAbstraction) DownloadFileAsString(path string) (string, error) {
	//nolint:exhaustruct // struct will be filled during unmarshal
	result := downloadFileAsStringResult{}
	err := testSetup.makeApiRequest("GET", "bfs/downloadFileAsString?path="+url.QueryEscape(path), &result, url.Values{})
	if err != nil {
		return "", err
	}
	return result.Content, nil
}

// DownloadFile downloads a file from the default BucketFS bucket to a local file.
func (testSetup TestSetupAbstraction) DownloadFile(remotePath string, localPath string) error {
	return testSetup.makeApiRequest("GET", "bfs/downloadFile?remotePath="+url.QueryEscape(remotePath)+"&localPath="+url.QueryEscape(localPath), &successResult{Success: true}, url.Values{})
}

// DeleteFile deletes a file from the default BucketFS bucket.
func (testSetup TestSetupAbstraction) DeleteFile(path string) error {
	return testSetup.makeApiRequest("DELETE", "bfs/deleteFile", &successResult{Success: true}, url.Values{
		"path": {path},
	})
}

// ListFiles lists files in the default BucketFS bucket.
func (testSetup TestSetupAbstraction) ListFiles(path string) ([]string, error) {
	//nolint:exhaustruct // struct will be filled during unmarshal
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

// Address of a service with hostname and port.
type ServiceAddress struct {
	HostName string `json:"hostName"`
	Port     int    `json:"port"`
}
