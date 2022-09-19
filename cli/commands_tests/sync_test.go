package commands_tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSync(t *testing.T) {
	i := newIntegrationTest(t)
	output, err := i.runCommand(t, "gen", "source", "test")
	fmt.Println(output)
	require.NoError(t, err)
	output, err = i.runCommand(t, "gen", "destination", "postgresql")
	fmt.Println(output)
	require.NoError(t, err)

	output, err = i.runCommand(t, "sync", ".")
	fmt.Println(output)
	require.NoError(t, err)
}
