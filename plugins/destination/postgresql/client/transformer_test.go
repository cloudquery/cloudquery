package client

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestStripNullsFromMarshalledJson(t *testing.T) {
	testcases := []struct {
		input          string
		expectedOutput string
	}{
		{
			input:          `{"hello":"world\u0000!","obj":{"nestedobj":{"num":"2500\u0000!","other":"bla"},"num":"1500"}}`,
			expectedOutput: `{"hello":"world!","obj":{"nestedobj":{"num":"2500!","other":"bla"},"num":"1500"}}`,
		},
		{
			input:          `"hello\u0000world!"`,
			expectedOutput: `"helloworld!"`,
		},
		{
			input:          `[{"SomeObj":"WithNull\u0000!"},{"A": "B","nested":{"property":"obj\u0000!"}}]`,
			expectedOutput: `[{"SomeObj":"WithNull!"},{"A":"B","nested":{"property":"obj!"}}]`,
		},
	}

	for _, testcase := range testcases {
		output := stripNullsFromMarshalledJson([]byte(testcase.input))
		if !slices.Equal(output, []byte(testcase.expectedOutput)) {
			t.Errorf("Expected output to be %s, but got %s", testcase.expectedOutput, output)
		}
	}
}
