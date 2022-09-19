package commands_tests

import (
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/Masterminds/semver/v3"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

type plugin struct {
	kind string
	name string
}

var pluginsToTest = []plugin{{kind: "source", name: "test"}, {kind: "destination", name: "postgresql"}}

type spec struct {
	Version string `yaml:"version"`
}

type config struct {
	Kind string `yaml:"kind"`
	Spec spec   `yaml:"spec"`
}

func validateYaml(t *testing.T, file string) config {
	configBytes, err := os.ReadFile(file)
	require.NoError(t, err)
	config := config{}
	err = yaml.Unmarshal(configBytes, &config)
	require.NoError(t, err)

	return config
}

func TestGen(t *testing.T) {
	for _, plugin := range pluginsToTest {
		t.Run(fmt.Sprintf("gen %s %s", plugin.kind, plugin.name), func(t *testing.T) {
			i := newIntegrationTest(t)
			output, err := i.runCommand(t, "gen", plugin.kind, plugin.name)
			fmt.Println(output)
			require.NoError(t, err)
			require.Contains(t, output, "plugin config successfully written to "+plugin.name+".yml")
			config := validateYaml(t, path.Join(i.dir, plugin.name+".yml"))
			// sanity check only; we shouldn't be testing the full config here
			require.Equal(t, plugin.kind, config.Kind)

			// remove `v` prefix from version to pass the strict parsing
			_, err = semver.StrictNewVersion(config.Spec.Version[1:])
			require.NoError(t, err)
		})
	}
}
