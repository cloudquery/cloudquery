package cmd

import (
	"testing"
)

func TestSync(t *testing.T) {
	// this works but some funny stuff is going on
	// _, filename, _, _ := runtime.Caller(0)
	// currentDir := path.Dir(filename)
	// testDataDir := path.Join(currentDir, "testdata")
	// cmd := newCmdRoot()
	// cmd.SetArgs([]string{"sync", testDataDir})
	// if err := cmd.Execute(); err != nil {
	// 	t.Fatal(err)
	// }
}
