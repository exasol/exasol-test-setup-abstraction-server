package exasol_test_setup_abstraction_go

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/suite"
)

// Default version specified in exasol-test-setup-abstraction-java.
const DEFAULT_EXASOL_VERSION = "7.1.23"
const NON_DEFAULT_EXASOL_VERSION = "8.23.0"

type BuilderSuite struct {
	suite.Suite
	setup *TestSetupAbstraction
}

func TestBuilderSuite(t *testing.T) {
	suite.Run(t, new(BuilderSuite))
}

func (suite *BuilderSuite) BeforeTest(suiteName, testName string) {
	suite.setup = nil
}

func (suite *BuilderSuite) AfterTest(suiteName, testName string) {
	if suite.setup != nil {
		suite.NoError(suite.setup.Stop())
	}
}

func (suite *BuilderSuite) TestDefaultConfiguration() {
	var err error
	suite.setup, err = New().Start()
	suite.NoError(err)
	suite.Equal(DEFAULT_EXASOL_VERSION, suite.getExasolDbVersion())
}

func (suite *BuilderSuite) TestCustomMissingConfigFile() {
	var err error
	suite.setup, err = New().CloudSetupConfigFilePath("missing-config-file.json").Start()
	suite.NoError(err)
	suite.Equal(DEFAULT_EXASOL_VERSION, suite.getExasolDbVersion())
}

func (suite *BuilderSuite) TestConfigFileWithInvalidContent() {
	var err error
	suite.setup, err = New().CloudSetupConfigFilePath(suite.writeTempFile("invalid json content")).Start()
	suite.ErrorContains(err, "failed to start server. Server did not print a port number")
	suite.ErrorContains(err, "E-ETSAS-7: Failed to start server: 'Unexpected char 105 at")
	suite.Nil(suite.setup)
}

func (suite *BuilderSuite) TestCustomExasolVersion() {
	var err error
	suite.setup, err = New().DockerDbVersion(NON_DEFAULT_EXASOL_VERSION).Start()
	suite.NoError(err)
	suite.Equal(NON_DEFAULT_EXASOL_VERSION, suite.getExasolDbVersion())
}

func (suite *BuilderSuite) writeTempFile(content string) string {
	tempDir := suite.T().TempDir()
	file := path.Join(tempDir, "temp-file")
	err := os.WriteFile(file, []byte(content), 0600)
	suite.NoError(err)
	return file
}

func (suite *BuilderSuite) getExasolDbVersion() string {
	db, err := suite.setup.CreateConnection()
	suite.NoError(err)
	defer db.Close()
	row := db.QueryRow("select param_value from exa_metadata where param_name = 'databaseProductVersion'")
	suite.NoError(err)
	var result string
	suite.NoError(row.Scan(&result))
	return result
}
