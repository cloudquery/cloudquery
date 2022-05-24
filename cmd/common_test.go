package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type CommandTestCases struct {
	Name           string
	PreCommands    []CommandTestCases
	Command        string
	ExpectedOutput string
	ExpectError    bool
	Args           []string
}

func testCommand(t *testing.T, tc CommandTestCases) {
	if tc.PreCommands != nil {
		for _, p := range tc.PreCommands {
			testCommand(t, p)
		}
	}
	args := append([]string{tc.Command}, tc.Args...)
	rootCmd.SetArgs(args)
	err := rootCmd.Execute()
	if !tc.ExpectError {
		require.NoError(t, err,
			"Input args: %v\nExpect out: %v\n", tc.Args, tc.ExpectedOutput)
	} else {
		assert.NotNil(t, err)
	}
}
