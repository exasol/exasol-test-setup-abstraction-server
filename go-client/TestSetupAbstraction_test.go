package exasol_test_setup_abstraction_go

import (
	"os"
	"testing"

	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/suite"
)

type TestSetupAbstractionSuite struct {
	suite.Suite
}

func TestTestSetupAbstractionSuite(t *testing.T) {
	suite.Run(t, new(TestSetupAbstractionSuite))
}

func (suite *TestSetupAbstractionSuite) TestCreateAndStop() {
	testSetup := Create("nonExistingConfig.json")
	testSetup.Stop()
}

func (suite *TestSetupAbstractionSuite) TestCreateConnection() {
	testSetup := Create("nonExistingConfig.json")
	defer testSetup.Stop()
	connection := testSetup.CreateConnection()
	defer func() { suite.NoError(connection.Close()) }()
	row := connection.QueryRow("SELECT 1")
	var result int
	suite.NoError(row.Scan(&result))
	suite.Assert().Equal(1, result)
}

func (suite *TestSetupAbstractionSuite) TestMakeDatabaseTcpServiceAccessibleFromLocalhost() {
	testSetup := Create("nonExistingConfig.json")
	defer testSetup.Stop()
	ports := testSetup.MakeDatabaseTcpServiceAccessibleFromLocalhost(123)
	suite.Assert().Equal(1, len(ports))
}

func (suite *TestSetupAbstractionSuite) TestMakeLocalTcpServiceAccessibleFromDatabase() {
	testSetup := Create("nonExistingConfig.json")
	defer testSetup.Stop()
	serviceAddress := testSetup.MakeLocalTcpServiceAccessibleFromDatabase(123)
	suite.Assert().NotNil(serviceAddress)
	suite.Assert().NotEmpty(serviceAddress)
}

func (suite *TestSetupAbstractionSuite) TestMakeTcpServiceAccessibleFromDatabase() {
	testSetup := Create("nonExistingConfig.json")
	defer testSetup.Stop()
	serviceAddress := testSetup.MakeTcpServiceAccessibleFromDatabase(ServiceAddress{"localhost", 134})
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
