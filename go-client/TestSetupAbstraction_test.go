package exasol_test_setup_abstraction_go

import (
	"database/sql"
	"os"
	"path"
	"testing"

	"github.com/exasol/exasol-driver-go"
	"github.com/stretchr/testify/suite"
)

type TestSetupAbstractionSuite struct {
	suite.Suite
	testSetup TestSetupAbstraction
}

func TestTestSetupAbstractionSuite(t *testing.T) {
	suite.Run(t, new(TestSetupAbstractionSuite))
}

func (suite *TestSetupAbstractionSuite) SetupSuite() {
	testSetup, err := New().Start()
	if err != nil {
		suite.FailNowf("failed to create test setup. Cause: %v", err.Error())
	}
	suite.testSetup = *testSetup
}

func (suite *TestSetupAbstractionSuite) TearDownSuite() {
	err := suite.testSetup.Stop()
	suite.Require().NoError(err)
}

func (suite *TestSetupAbstractionSuite) TestCreateConnection() {
	connection, err := suite.testSetup.CreateConnection()
	suite.Require().NoError(err)
	defer func() { suite.Require().NoError(connection.Close()) }()
	row := connection.QueryRow("SELECT 1")
	var result int
	suite.Require().NoError(row.Scan(&result))
	suite.Equal(1, result)
}

func (suite *TestSetupAbstractionSuite) TestCreateConnectionWithAutocommitOff() {
	connection, err := suite.testSetup.CreateConnectionWithConfig(false)
	suite.Require().NoError(err)
	defer func() { suite.Require().NoError(connection.Close()) }()
	row := connection.QueryRow("SELECT 1")
	var result int
	suite.Require().NoError(row.Scan(&result))
	suite.Equal(1, result)
}

func (suite *TestSetupAbstractionSuite) TestGetConnectionInfo() {
	info, err := suite.testSetup.GetConnectionInfo()
	suite.Require().NoError(err)
	suite.Equal("SYS", info.User)
	suite.Equal("exasol", info.Password)
	suite.NotEmpty(info.Host)
	suite.Positive(info.Port)
}

func (suite *TestSetupAbstractionSuite) TestGetConnectionInfoReturnsValidSettings() {
	info, err := suite.testSetup.GetConnectionInfo()
	suite.Require().NoError(err)
	db, err := sql.Open("exasol", exasol.NewConfig(info.User, info.Password).
		Port(info.Port).
		Host(info.Host).
		ValidateServerCertificate(false).
		String())
	suite.Require().NoError(err)
	defer func() { suite.Require().NoError(db.Close()) }()
	row := db.QueryRow("SELECT 1")
	var result int
	suite.Require().NoError(row.Scan(&result))
	suite.Equal(1, result)
}

func (suite *TestSetupAbstractionSuite) TestMakeDatabaseTcpServiceAccessibleFromLocalhost() {
	ports, err := suite.testSetup.MakeDatabaseTcpServiceAccessibleFromLocalhost(123)
	suite.Require().NoError(err)
	suite.Len(ports, 1)
}

func (suite *TestSetupAbstractionSuite) TestMakeDatabaseTcpServiceAccessibleFromLocalhostFails() {
	_, err := suite.testSetup.MakeDatabaseTcpServiceAccessibleFromLocalhost(-123)
	suite.Require().ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-ETSAS-8: Port number -123 is negative. Please specify a valid port.\"")
}

func (suite *TestSetupAbstractionSuite) TestMakeLocalTcpServiceAccessibleFromDatabase() {
	serviceAddress, err := suite.testSetup.MakeLocalTcpServiceAccessibleFromDatabase(123)
	suite.Require().NoError(err)
	suite.NotNil(serviceAddress)
	suite.NotEmpty(serviceAddress)
}

func (suite *TestSetupAbstractionSuite) TestMakeLocalTcpServiceAccessibleFromDatabaseFails() {
	_, err := suite.testSetup.MakeLocalTcpServiceAccessibleFromDatabase(-123)
	suite.Require().ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-ETSAS-8: Port number -123 is negative. Please specify a valid port.\"")
}

func (suite *TestSetupAbstractionSuite) TestMakeTcpServiceAccessibleFromDatabase() {
	serviceAddress, err := suite.testSetup.MakeTcpServiceAccessibleFromDatabase(ServiceAddress{"localhost", 134})
	suite.Require().NoError(err)
	suite.NotNil(serviceAddress)
	suite.NotEmpty(serviceAddress)
}

func (suite *TestSetupAbstractionSuite) TestMakeTcpServiceAccessibleFromDatabaseFails() {
	_, err := suite.testSetup.MakeTcpServiceAccessibleFromDatabase(ServiceAddress{"localhost", -123})
	suite.Require().ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-ETSAS-8: Port number -123 is negative. Please specify a valid port.\"")
}

func (suite *TestSetupAbstractionSuite) TestUploadStringContent() {
	err := suite.testSetup.UploadStringContent("test", "TestUploadStringContent.txt")
	suite.Require().NoError(err)
	defer func() { suite.Require().NoError(suite.testSetup.DeleteFile("TestUploadStringContent.txt")) }()
	files, err := suite.testSetup.ListFiles("")
	suite.Require().NoError(err)
	suite.Contains(files, "TestUploadStringContent.txt")
}

func (suite *TestSetupAbstractionSuite) TestDownloadStringContent() {
	err := suite.testSetup.UploadStringContent("test", "TestDownloadStringContent.txt")
	suite.Require().NoError(err)
	defer func() { suite.Require().NoError(suite.testSetup.DeleteFile("TestDownloadStringContent.txt")) }()
	content, err := suite.testSetup.DownloadFileAsString("TestDownloadStringContent.txt")
	suite.Require().NoError(err)
	suite.Equal("test", content)
}

func (suite *TestSetupAbstractionSuite) TestDownloadFile() {
	err := suite.testSetup.UploadStringContent("test", "TestDownloadFile.txt")
	suite.Require().NoError(err)
	defer func() { suite.Require().NoError(suite.testSetup.DeleteFile("TestDownloadFile.txt")) }()
	tempDir := suite.T().TempDir()
	targetFile := path.Join(tempDir, "myFile.txt")
	err = suite.testSetup.DownloadFile("TestDownloadFile.txt", targetFile)
	suite.Require().NoError(err)
	content, err := os.ReadFile(targetFile)
	suite.Require().NoError(err)
	suite.Equal("test", string(content))
}

func (suite *TestSetupAbstractionSuite) TestDownloadNonExistingFileFails() {
	tempDir := suite.T().TempDir()
	targetFile := path.Join(tempDir, "myFile.txt")
	err := suite.testSetup.DownloadFile("non-existing-file.txt", targetFile)
	suite.Require().ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-BFSJ-2: File or directory not found trying to download 'https://")
}

func (suite *TestSetupAbstractionSuite) TestDeleteFile() {
	err := suite.testSetup.UploadStringContent("test", "TestDeleteFile.txt")
	suite.Require().NoError(err)
	err = suite.testSetup.DeleteFile("TestDeleteFile.txt")
	suite.Require().NoError(err)
	files, err := suite.testSetup.ListFiles("")
	suite.Require().NoError(err)
	suite.NotContains(files, "TestDeleteFile.txt")
}

func (suite *TestSetupAbstractionSuite) TestDeleteNonExistingFileSucceeds() {
	err := suite.testSetup.DeleteFile("non-existing-file.txt")
	suite.Require().NoError(err)
}

func (suite *TestSetupAbstractionSuite) TestUploadFile() {
	tmpFile := suite.createTestFile()
	err := suite.testSetup.UploadFile(tmpFile, "TestUploadFile.txt")
	suite.Require().NoError(err)
	defer func() { suite.Require().NoError(suite.testSetup.DeleteFile("TestUploadFile.txt")) }()
	files, err := suite.testSetup.ListFiles("")
	suite.Require().NoError(err)
	suite.Contains(files, "TestUploadFile.txt")
}

func (suite *TestSetupAbstractionSuite) TestUploadNonExistingFileFails() {
	err := suite.testSetup.UploadFile("non-existing-local-file.txt", "TestUploadFile.txt")
	suite.Require().ErrorContains(err, "non-existing-local-file.txt not found")
}

func (suite *TestSetupAbstractionSuite) TestListFilesInRootDir() {
	files, err := suite.testSetup.ListFiles("/")
	suite.Require().NoError(err)
	suite.Len(files, 1)
	suite.Contains(files, "") // BucketFS in Exasol 8 does not contain "EXAClusterOS/" any more
}

//nolint:testifylint // We have three assertions for the same error
func (suite *TestSetupAbstractionSuite) TestListFilesNonExistingDirectory() {
	_, err := suite.testSetup.ListFiles("non-existing-dir")
	suite.ErrorContains(err, "request failed with status 500 (500 Server Error).")
	suite.ErrorContains(err, "E-BFSJ-11: Unable to list contents of 'non-existing-dir' in bucket 'http")
	suite.ErrorContains(err, "No such file or directory.\"")
}

func (suite *TestSetupAbstractionSuite) createTestFile() string {
	tempDir := suite.T().TempDir()
	file, err := os.CreateTemp(tempDir, "my-temp-file-*")
	if err != nil {
		panic(err)
	}
	defer func() { suite.Require().NoError(file.Close()) }()
	_, err = file.WriteString("test")
	if err != nil {
		panic(err)
	}
	return file.Name()
}
