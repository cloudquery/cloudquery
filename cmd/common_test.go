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
	Args           []string
}

func testCommand(t *testing.T, testCases []CommandTestCases) {
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			if tc.PreCommands != nil {
				testCommand(t, tc.PreCommands)
			}
			args := append([]string{tc.Command}, tc.Args...)
			rootCmd.SetArgs(args)
			out := capturer.CaptureOutput(func() {
				err := rootCmd.Execute()

				require.NoError(t, err,
					"Input args: %v\nExpect out: %v\n", tc.Args, tc.ExpectedOutput)
			})
			assert.Contains(t, out, tc.ExpectedOutput,
				"Expect: %v\nActual: %v\n", tc.ExpectedOutput, out)
		})
	}
}
