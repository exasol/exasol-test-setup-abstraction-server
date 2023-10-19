package exasol_test_setup_abstraction_go

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"regexp"
	"strconv"
	"sync"
	"time"
)

type serverProcess struct {
	process        *exec.Cmd
	serverEndpoint string
	stopped        *bool
	stoppedMutex   *sync.Mutex
	errorStream    *bytes.Buffer
}

// startServer starts the server in the given version and with the given config file.
func startServer(serverVersion string, config Builder) (*serverProcess, error) {
	serverPath, err := downloadServerIfNotPresent()
	if err != nil {
		return nil, err
	}
	args := getServerProcessArguments(serverPath, config)
	log.Printf("Starting server version %s with arguments %v", serverVersion, args)
	process := exec.Command("java", args...)
	var outputStream, errorStream bytes.Buffer
	process.Stdout = &outputStream
	process.Stderr = &errorStream
	stoppedMutex := &sync.Mutex{}
	stopped := false
	go waitForServer(process, &errorStream, &outputStream, &stopped, stoppedMutex)
	port, err := getServerPort(config.startupTimeout, &stopped, &outputStream, &errorStream)
	if err != nil {
		return nil, err
	}
	serverEndpoint := fmt.Sprintf("http://localhost:%v/", port)
	return &serverProcess{process: process, serverEndpoint: serverEndpoint, stopped: &stopped, stoppedMutex: stoppedMutex, errorStream: &errorStream}, nil
}

func getServerProcessArguments(serverPath string, config Builder) []string {
	var args = []string{}
	if config.dockerDbVersion != "" {
		args = append(args, "-Dcom.exasol.dockerdb.image="+config.dockerDbVersion)
	}
	args = append(args, "-jar", serverPath)
	if config.configFilePath != "" {
		args = append(args, config.configFilePath)
	}
	return args
}

func downloadServerIfNotPresent() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("failed to get home dir. Cause: %w", err)
	}
	serverDir := path.Join(homeDir, ".test-setup-abstraction-server")
	if _, err := os.Stat(serverDir); os.IsNotExist(err) {
		err := os.Mkdir(serverDir, 0700)
		if err != nil {
			return "", fmt.Errorf("failed to create directory %q: %w", serverDir, err)
		}
	}
	const serverJar = "exasol-test-setup-abstraction-server-" + serverVersion + ".jar"
	localPath := path.Join(serverDir, serverJar)
	url := "https://github.com/exasol/exasol-test-setup-abstraction-server/releases/download/" + serverVersion + "/" + serverJar
	if _, err := os.Stat(localPath); os.IsNotExist(err) {
		err = downloadFile(url, localPath)
		if err != nil {
			return "", err
		}
	}
	return localPath, nil
}

func getServerPort(timeout time.Duration, stopped *bool, output *bytes.Buffer, errorStream *bytes.Buffer) (int, error) {
	startTime := time.Now()
	pattern := regexp.MustCompile("Server running on port: (\\d+)\n")
	for {
		result := pattern.FindSubmatch(output.Bytes())
		if len(result) != 0 {
			portString := string(result[1])
			port, err := strconv.ParseInt(portString, 10, 32)
			if err != nil {
				return -1, fmt.Errorf("failed to parse port %q", portString)
			}
			return int(port), nil
		}
		duration := time.Since(startTime)
		if duration >= timeout {
			return -1, fmt.Errorf("failed to start server. Server did not print a port number after %v seconds. Output: '%s', Error: '%s'", duration, output, errorStream)
		}
		if !*stopped {
			time.Sleep(1 * time.Second)
		} else {
			return -1, fmt.Errorf("server stopped after %v seconds. Output: '%s', Error: '%s'", duration, output, errorStream)
		}
	}
}

func waitForServer(serverProcess *exec.Cmd, errorStream *bytes.Buffer, outputStream *bytes.Buffer, stopped *bool, stoppedMutex *sync.Mutex) {
	err := serverProcess.Run()
	if err != nil && !isStopped(stopped, stoppedMutex) { // after we killed the thread we expect an error
		fmt.Printf("failed to start server: %v. Error output: '%s', output stream: '%s'\n", err, errorStream.String(), outputStream.String())
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

func (testSetup *serverProcess) stop() error {
	testSetup.stoppedMutex.Lock()
	*testSetup.stopped = true
	testSetup.stoppedMutex.Unlock()
	err := testSetup.process.Process.Signal(os.Kill)
	if err != nil {
		return fmt.Errorf("failed to stop test-setup-abstraction server. Cause: %w", err)
	}
	return nil
}
