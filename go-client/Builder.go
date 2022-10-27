package exasol_test_setup_abstraction_go

type Builder struct {
	configFilePath  string
	dockerDbVersion string
}

// New creates a new builder that allows creating a new TestSetupAbstraction.
func New() Builder {
	return Builder{}
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
	c.dockerDbVersion = dockerDbVersion
	return c
}

// Start launches the test setup using the given configuration.
// Don't forget to stop the setup after usage by calling TestSetupAbstraction.Stop().
func (c Builder) Start() (*TestSetupAbstraction, error) {
	return startTestSetupAbstraction(c)
}
