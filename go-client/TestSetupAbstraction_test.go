package exasol_test_setup_abstraction_go

import (
	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/suite"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

type TestSetupAbstractionSuite struct {
	suite.Suite
	testSetup TestSetupAbstraction
}

func TestTestSetupAbstractionSuite(t *testing.T) {
	suite.Run(t, new(TestSetupAbstractionSuite))
}

func (suite *TestSetupAbstractionSuite) SetupSuite() {
	suite.testSetup = Create("nonExistingConfig.json")
}

func (suite *TestSetupAbstractionSuite) TearDownSuite() {
	suite.testSetup.Stop()
}

func (suite *TestSetupAbstractionSuite) TestCreateConnection() {
	connection := suite.testSetup.CreateConnection()
	defer func() { suite.NoError(connection.Close()) }()
	row := connection.QueryRow("SELECT 1")
	var result int
	suite.NoError(row.Scan(&result))
	suite.Assert().Equal(1, result)
}

func (suite *TestSetupAbstractionSuite) TestMakeDatabaseTcpServiceAccessibleFromLocalhost() {
	ports := suite.testSetup.MakeDatabaseTcpServiceAccessibleFromLocalhost(123)
	suite.Assert().Equal(1, len(ports))
}

func (suite *TestSetupAbstractionSuite) TestMakeLocalTcpServiceAccessibleFromDatabase() {
	serviceAddress := suite.testSetup.MakeLocalTcpServiceAccessibleFromDatabase(123)
	suite.Assert().NotNil(serviceAddress)
	suite.Assert().NotEmpty(serviceAddress)
}

func (suite *TestSetupAbstractionSuite) TestMakeTcpServiceAccessibleFromDatabase() {
	serviceAddress := suite.testSetup.MakeTcpServiceAccessibleFromDatabase(ServiceAddress{"localhost", 134})
	suite.Assert().NotNil(serviceAddress)
	suite.Assert().NotEmpty(serviceAddress)
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
	suite.testSetup.UploadStringContent("test", "TestUploadStringContent.txt")
	defer suite.testSetup.DeleteFile("TestUploadStringContent.txt")
	suite.Assert().Contains(suite.testSetup.ListFiles(""), "TestUploadStringContent.txt")
}

func (suite *TestSetupAbstractionSuite) TestDownloadStringContent() {
	suite.testSetup.UploadStringContent("test", "TestDownloadStringContent.txt")
	defer suite.testSetup.DeleteFile("TestDownloadStringContent.txt")
	suite.Assert().Equal("test", suite.testSetup.DownloadFileAsString("TestDownloadStringContent.txt"))
}

func (suite *TestSetupAbstractionSuite) TestDownloadFile() {
	defer suite.testSetup.printServerErrors()
	suite.testSetup.UploadStringContent("test", "TestDownloadFile.txt")
	defer suite.testSetup.DeleteFile("TestDownloadFile.txt")
	tempDir := createTempDir()
	defer suite.deleteFileOrFolder(tempDir)
	targetFile := path.Join(tempDir, "myFile.txt")
	suite.testSetup.DownloadFile("TestDownloadFile.txt", targetFile)
	content, err := ioutil.ReadFile(targetFile)
	if err != nil {
		panic(err)
	}
	suite.Assert().Equal("test", string(content))
}

func createTempDir() string {
	tempDir, err := os.MkdirTemp(os.TempDir(), "TestDownloadFile")
	if err != nil {
		panic(err)
	}
	return tempDir
}

func (suite *TestSetupAbstractionSuite) TestDeleteFile() {
	suite.testSetup.UploadStringContent("test", "TestDeleteFile.txt")
	suite.testSetup.DeleteFile("TestDeleteFile.txt")
	suite.Assert().NotContains(suite.testSetup.ListFiles(""), "TestDeleteFile.txt")
}

func (suite *TestSetupAbstractionSuite) TestUploadFile() {
	defer suite.testSetup.printServerErrors()
	tmpFile := suite.createTestFile()
	defer suite.deleteFileOrFolder(tmpFile)
	suite.testSetup.UploadFile(tmpFile, "TestUploadFile.txt")
	defer suite.testSetup.DeleteFile("TestUploadFile.txt")
	suite.Assert().Contains(suite.testSetup.ListFiles(""), "TestUploadFile.txt")
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
