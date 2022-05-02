package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
	"github.com/zenizh/go-capturer"
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
	out := capturer.CaptureOutput(func() {
		err := rootCmd.Execute()
		if !tc.ExpectError {
			require.NoError(t, err,
				"Input args: %v\nExpect out: %v\n", tc.Args, tc.ExpectedOutput)
		} else {
			assert.NotNil(t, err)
		}
	})
	require.Contains(t, out, tc.ExpectedOutput,
		"Expect: %v\nActual: %v\n", tc.ExpectedOutput, out)
}
