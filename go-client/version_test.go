package exasol_test_setup_abstraction_go

import (
	"os"
	"testing"

	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/suite"
	"gopkg.in/yaml.v3"
)

type VersionSuite struct {
	suite.Suite
}

func TestVersionTestSuite(t *testing.T) {
	suite.Run(t, new(VersionSuite))
}

func (suite *VersionSuite) TestServerVersionSameAsPom() {
	suite.Assert().Equal(suite.readVersionFromPom(), serverVersion, "version in pom.xml and constant match")
}

func (suite *VersionSuite) TestServerVersionSameAsProjectKeeperConfig() {
	suite.Assert().Equal(suite.readProjectKeeperConf().Version, serverVersion, "version in .project-keeper.yml and constant match")
}

func (suite *VersionSuite) readVersionFromPom() string {
	pomFile, err := os.Open("../server/pom.xml")
	suite.NoError(err)
	defer pomFile.Close()
	pom, err := xmlquery.Parse(pomFile)
	suite.NoError(err)
	return xmlquery.FindOne(pom, "/project/version").InnerText()
}

type projectKeeperConfig struct {
	Version string `yaml:"version"`
}

func (suite *VersionSuite) readProjectKeeperConf() *projectKeeperConfig {
	yamlFile, err := os.ReadFile("../.project-keeper.yml")
	if err != nil {
		suite.FailNowf("failed to read project keeper file: %v", err.Error())
	}
	config := &projectKeeperConfig{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		suite.FailNowf("unmarshal failed: %v", err.Error())
	}
	return config
}
