package exasol_test_setup_abstraction_go

type Configuration struct {
	configFilePath  string
	dockerDbVersion string
}

// New creates a new configuration object that allows creating a new TestSetupAbstraction.
func New() Configuration {
	return Configuration{}
}

// CloudSetupConfigFilePath sets the path to the cloud setup config file.
// Call this to use a cloud test setup e.g. in AWS.
// This will fall back to a local Docker container in case the file does not exist.
func (c Configuration) CloudSetupConfigFilePath(path string) Configuration {
	c.configFilePath = path
	return c
}

// DockerDbVersion sets the Exasol Docker DB version to start.
// This defaults to the version defined in exasol-test-setup-abstraction-java.
func (c Configuration) DockerDbVersion(dockerDbVersion string) Configuration {
	c.dockerDbVersion = dockerDbVersion
	return c
}

// Start launches the test setup using the given configuration.
// Don't forget to stop the setup after usage by calling TestSetupAbstraction.Stop().
func (c Configuration) Start() (*TestSetupAbstraction, error) {
	return startTestSetupAbstraction(c)
}
