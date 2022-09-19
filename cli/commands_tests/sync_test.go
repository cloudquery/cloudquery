package commands_tests

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/require"
)

func addDestination(t *testing.T, configFile string) {
	input, err := os.ReadFile(configFile)
	require.NoError(t, err)

	output := bytes.ReplaceAll(input, []byte("destinations: []"), []byte("destinations: [postgresql]"))
	err = os.WriteFile(configFile, output, 0666)
	require.NoError(t, err)
}

func TestSync(t *testing.T) {
	i := newIntegrationTest(t)
	output, err := i.runCommand(t, "gen", "source", "test")
	fmt.Println(output)
	require.NoError(t, err)
	output, err = i.runCommand(t, "gen", "destination", "postgresql")
	fmt.Println(output)
	require.NoError(t, err)

	addDestination(t, path.Join(i.dir, "test.yml"))

	output, err = i.runCommand(t, "sync", ".")
	fmt.Println(output)
	require.NoError(t, err)
}
