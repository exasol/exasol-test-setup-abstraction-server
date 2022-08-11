package exasol_test_setup_abstraction_go

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/antchfx/xmlquery"
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
	testSetup, err := Create("nonExistingConfig.json")
	if err != nil {
		panic(fmt.Sprintf("failed to create test setup. Cause: %v", err.Error()))
	}
	suite.testSetup = *testSetup
}

func (suite *TestSetupAbstractionSuite) TearDownSuite() {
	err := suite.testSetup.Stop()
	suite.NoError(err)
}

func (suite *TestSetupAbstractionSuite) TestCreateConnection() {
	connection, err := suite.testSetup.CreateConnection()
	suite.NoError(err)
	defer func() { suite.NoError(connection.Close()) }()
	row := connection.QueryRow("SELECT 1")
	var result int
	suite.NoError(row.Scan(&result))
	suite.Assert().Equal(1, result)
}

func (suite *TestSetupAbstractionSuite) TestCreateConnectionWithAutocommitOff() {
	connection, err := suite.testSetup.CreateConnectionWithConfig(false)
	suite.NoError(err)
	defer func() { suite.NoError(connection.Close()) }()
	row := connection.QueryRow("SELECT 1")
	var result int
	suite.NoError(row.Scan(&result))
	suite.Assert().Equal(1, result)
}

func (suite *TestSetupAbstractionSuite) TestGetConnectionInfo() {
	info, err := suite.testSetup.GetConnectionInfo()
	suite.NoError(err)
	suite.Equal("SYS", info.User)
	suite.Equal("exasol", info.Password)
	suite.NotEmpty(info.Host)
	suite.Greater(info.Port, 0)
}

func (suite *TestSetupAbstractionSuite) TestGetConnectionInfoReturnsValidSettings() {
	info, err := suite.testSetup.GetConnectionInfo()
	suite.NoError(err)
	db, err := sql.Open("exasol", exasol.NewConfig(info.User, info.Password).
		Port(info.Port).
		Host(info.Host).
		ValidateServerCertificate(false).
		String())
	suite.NoError(err)
	defer func() { suite.NoError(db.Close()) }()
	row := db.QueryRow("SELECT 1")
	var result int
	suite.NoError(row.Scan(&result))
	suite.Assert().Equal(1, result)
}

func (suite *TestSetupAbstractionSuite) TestMakeDatabaseTcpServiceAccessibleFromLocalhost() {
	ports, err := suite.testSetup.MakeDatabaseTcpServiceAccessibleFromLocalhost(123)
	suite.NoError(err)
	suite.Assert().Equal(1, len(ports))
}

func (suite *TestSetupAbstractionSuite) TestMakeDatabaseTcpServiceAccessibleFromLocalhostFails() {
	_, err := suite.testSetup.MakeDatabaseTcpServiceAccessibleFromLocalhost(-123)
	suite.ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-ETSAS-8: Port number -123 is negative. Please specify a valid port.\"")
}

func (suite *TestSetupAbstractionSuite) TestMakeLocalTcpServiceAccessibleFromDatabase() {
	serviceAddress, err := suite.testSetup.MakeLocalTcpServiceAccessibleFromDatabase(123)
	suite.NoError(err)
	suite.Assert().NotNil(serviceAddress)
	suite.Assert().NotEmpty(serviceAddress)
}

func (suite *TestSetupAbstractionSuite) TestMakeLocalTcpServiceAccessibleFromDatabaseFails() {
	_, err := suite.testSetup.MakeLocalTcpServiceAccessibleFromDatabase(-123)
	suite.ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-ETSAS-8: Port number -123 is negative. Please specify a valid port.\"")
}

func (suite *TestSetupAbstractionSuite) TestMakeTcpServiceAccessibleFromDatabase() {
	serviceAddress, err := suite.testSetup.MakeTcpServiceAccessibleFromDatabase(ServiceAddress{"localhost", 134})
	suite.NoError(err)
	suite.Assert().NotNil(serviceAddress)
	suite.Assert().NotEmpty(serviceAddress)
}

func (suite *TestSetupAbstractionSuite) TestMakeTcpServiceAccessibleFromDatabaseFails() {
	_, err := suite.testSetup.MakeTcpServiceAccessibleFromDatabase(ServiceAddress{"localhost", -123})
	suite.ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-ETSAS-8: Port number -123 is negative. Please specify a valid port.\"")
}

func (suite *TestSetupAbstractionSuite) TestServerVersionCorrect() {
	suite.Assert().Equal(suite.readVersionFromPom(), serverVersion, "version in pom.xml and constant match")
}

func (suite *TestSetupAbstractionSuite) readVersionFromPom() string {
	pomFile, err := os.Open("../server/pom.xml")
	suite.NoError(err)
	defer pomFile.Close()
	pom, err := xmlquery.Parse(pomFile)
	suite.NoError(err)
	return xmlquery.FindOne(pom, "/project/version").InnerText()
}

func (suite *TestSetupAbstractionSuite) TestUploadStringContent() {
	err := suite.testSetup.UploadStringContent("test", "TestUploadStringContent.txt")
	suite.NoError(err)
	defer func() { suite.NoError(suite.testSetup.DeleteFile("TestUploadStringContent.txt")) }()
	files, err := suite.testSetup.ListFiles("")
	suite.NoError(err)
	suite.Assert().Contains(files, "TestUploadStringContent.txt")
}

func (suite *TestSetupAbstractionSuite) TestDownloadStringContent() {
	err := suite.testSetup.UploadStringContent("test", "TestDownloadStringContent.txt")
	suite.NoError(err)
	defer func() { suite.NoError(suite.testSetup.DeleteFile("TestDownloadStringContent.txt")) }()
	content, err := suite.testSetup.DownloadFileAsString("TestDownloadStringContent.txt")
	suite.NoError(err)
	suite.Assert().Equal("test", content)
}

func (suite *TestSetupAbstractionSuite) TestDownloadFile() {
	err := suite.testSetup.UploadStringContent("test", "TestDownloadFile.txt")
	suite.NoError(err)
	defer func() { suite.NoError(suite.testSetup.DeleteFile("TestDownloadFile.txt")) }()
	tempDir := createTempDir()
	defer suite.deleteFileOrFolder(tempDir)
	targetFile := path.Join(tempDir, "myFile.txt")
	err = suite.testSetup.DownloadFile("TestDownloadFile.txt", targetFile)
	suite.NoError(err)
	content, err := ioutil.ReadFile(targetFile)
	suite.NoError(err)
	suite.Assert().Equal("test", string(content))
}

func (suite *TestSetupAbstractionSuite) TestDownloadNonExistingFileFails() {
	tempDir := createTempDir()
	defer suite.deleteFileOrFolder(tempDir)
	targetFile := path.Join(tempDir, "myFile.txt")
	err := suite.testSetup.DownloadFile("non-existing-file.txt", targetFile)
	suite.ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-BFSJ-2: File or directory not found trying to download http://")
}

func createTempDir() string {
	tempDir, err := os.MkdirTemp(os.TempDir(), "TestDownloadFile")
	if err != nil {
		panic(err)
	}
	return tempDir
}

func (suite *TestSetupAbstractionSuite) TestDeleteFile() {
	err := suite.testSetup.UploadStringContent("test", "TestDeleteFile.txt")
	suite.NoError(err)
	err = suite.testSetup.DeleteFile("TestDeleteFile.txt")
	suite.NoError(err)
	files, err := suite.testSetup.ListFiles("")
	suite.NoError(err)
	suite.Assert().NotContains(files, "TestDeleteFile.txt")
}

func (suite *TestSetupAbstractionSuite) TestDeleteNonExistingFileSucceeds() {
	err := suite.testSetup.DeleteFile("non-existing-file.txt")
	suite.NoError(err)
}

func (suite *TestSetupAbstractionSuite) TestUploadFile() {
	tmpFile := suite.createTestFile()
	defer suite.deleteFileOrFolder(tmpFile)
	err := suite.testSetup.UploadFile(tmpFile, "TestUploadFile.txt")
	suite.NoError(err)
	defer func() { suite.NoError(suite.testSetup.DeleteFile("TestUploadFile.txt")) }()
	files, err := suite.testSetup.ListFiles("")
	suite.NoError(err)
	suite.Assert().Contains(files, "TestUploadFile.txt")
}

func (suite *TestSetupAbstractionSuite) TestUploadNonExistingFileFails() {
	err := suite.testSetup.UploadFile("non-existing-local-file.txt", "TestUploadFile.txt")
	suite.ErrorContains(err, "non-existing-local-file.txt not found")
}

func (suite *TestSetupAbstractionSuite) TestListFilesInRootDir() {
	files, err := suite.testSetup.ListFiles("/")
	suite.NoError(err)
	suite.Equal(1, len(files))
	suite.Contains(files, "EXAClusterOS")
}

func (suite *TestSetupAbstractionSuite) TestListFilesNonExistingDirectory() {
	_, err := suite.testSetup.ListFiles("non-existing-dir")
	suite.ErrorContains(err, "request failed with status 500 (500 Server Error). Response: \"E-BFSJ-11: Unable to list contents of 'non-existing-dir' in bucket bfsdefault/default: No such file or directory.\"")
}

func (suite *TestSetupAbstractionSuite) createTestFile() string {
	file, err := ioutil.TempFile(os.TempDir(), "my-temp-file-*")
	if err != nil {
		panic(err)
	}
	defer func() { suite.NoError(file.Close()) }()
	_, err = file.WriteString("test")
	if err != nil {
		panic(err)
	}
	return file.Name()
}

func (suite *TestSetupAbstractionSuite) deleteFileOrFolder(file string) {
	err := os.RemoveAll(file)
	if err != nil {
		panic(err)
	}
}
