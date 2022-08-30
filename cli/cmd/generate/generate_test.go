package generate

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	cmd := NewCmdGenerate()
	cmd.SetArgs([]string{"source", "aws"})
	if err := cmd.Execute(); err != nil {
		t.Fatal(err)
	}
}
