package exasol_test_setup_abstraction_go

import "time"

type Builder struct {
	configFilePath  string
	exasolDbVersion string
	startupTimeout  time.Duration
}

// New creates a new builder that allows creating a new TestSetupAbstraction.
func New() Builder {
	return Builder{
		configFilePath:  "",
		exasolDbVersion: "8.23.1",
		startupTimeout:  time.Minute * 10,
	}
}

// CloudSetupConfigFilePath sets the path to the cloud setup config file.
// Call this to use a cloud test setup e.g. in AWS.
// This will fall back to a local Docker container in case the file does not exist.
func (c Builder) CloudSetupConfigFilePath(path string) Builder {
	c.configFilePath = path
	return c
}

// DockerDbVersion sets the Exasol Docker DB version to start.
// This defaults to the version defined in exasol-test-setup-abstraction-java.
func (c Builder) DockerDbVersion(dockerDbVersion string) Builder {
	c.exasolDbVersion = dockerDbVersion
	return c
}

// StartupTimeout sets the timeout for starting the Exasol test setup.
// This defaults to 10 minutes.
func (c Builder) StartupTimeout(timeout time.Duration) Builder {
	c.startupTimeout = timeout
	return c
}

// Start launches the test setup using the given configuration.
// Don't forget to stop the setup after usage by calling TestSetupAbstraction.Stop().
func (c Builder) Start() (*TestSetupAbstraction, error) {
	return startTestSetupAbstraction(c)
}
