package exasol_test_setup_abstraction_go

import (
	"os"
	"testing"

	"github.com/antchfx/xmlquery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

// These tests verify consistent project versions.

func TestServerVersionSameAsPom(t *testing.T) {
	assert.Equal(t, serverVersion, readVersionFromPom(t), "version in pom.xml and constant match")
}

func TestServerVersionSameAsProjectKeeperConfig(t *testing.T) {
	assert.Equal(t, serverVersion, readProjectKeeperConf(t).Version, "version in .project-keeper.yml and constant match")
}

func readVersionFromPom(t *testing.T) string {
	pomFile, err := os.Open("../server/pom.xml")
	require.NoError(t, err)
	defer pomFile.Close()
	pom, err := xmlquery.Parse(pomFile)
	require.NoError(t, err)
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
