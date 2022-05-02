package cmd

import (
	"testing"
)

func Test_Init(t *testing.T) {

	testCases := []CommandTestCases{
		{
			Name:           "init-no-args",
			Command:        "init",
			ExpectError:    true,
			ExpectedOutput: "Error: requires at least 1 arg(s), only received 0",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			testCommand(t, tc)
		})
	}
}
