package exasol_test_setup_abstraction_go

import (
	"os"
	"testing"

	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

// These tests verify consistent project versions.

func TestServerVersionSameAsPom(t *testing.T) {
	assert.Equal(t, readVersionFromPom(t), serverVersion, "version in pom.xml and constant match")
}

func TestServerVersionSameAsProjectKeeperConfig(t *testing.T) {
	assert.Equal(t, readProjectKeeperConf(t).Version, serverVersion, "version in .project-keeper.yml and constant match")
}

func readVersionFromPom(t *testing.T) string {
	pomFile, err := os.Open("../server/pom.xml")
	assert.NoError(t, err)
	defer pomFile.Close()
	pom, err := xmlquery.Parse(pomFile)
	assert.NoError(t, err)
	return xmlquery.FindOne(pom, "/project/version").InnerText()
}

type projectKeeperConfig struct {
	Version string `yaml:"version"`
}

func readProjectKeeperConf(t *testing.T) *projectKeeperConfig {
	yamlFile, err := os.ReadFile("../.project-keeper.yml")
	if err != nil {
		t.Fatalf("failed to read project keeper file: %v", err)
	}
	//nolint:exhaustruct // struct will be filled during unmarshal
	config := &projectKeeperConfig{}
	err = yaml.Unmarshal(yamlFile, config)
	if err != nil {
		t.Fatalf("Failed to parse YAML in project keeper configuration: %v", err.Error())
	}
	return config
}
