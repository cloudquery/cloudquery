package cmd

import (
	"testing"
)

func Test_Version(t *testing.T) {
	testCommand(t, []CommandTestCases{
		{
			Name:           "simple-version",
			Command:        "version",
			ExpectedOutput: "Version: development",
		},
	})
}
