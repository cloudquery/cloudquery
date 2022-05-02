package cmd

import (
	"testing"
)

func Test_Version(t *testing.T) {
	testCases := []CommandTestCases{
		{
			Name:           "simple-version",
			Command:        "version",
			ExpectedOutput: "Version: development",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testCommand(t, tc)
		})
	}
}
