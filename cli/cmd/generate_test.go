package cmd

import (
	"io/ioutil"
	"os"
	"path"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGenerate(t *testing.T) {
	tmpdir, tmpErr := os.MkdirTemp("", "generate_test_*")
	if tmpErr != nil {
		t.Fatalf("failed to create temporary directory: %v", tmpErr)
	}
	defer os.RemoveAll(tmpdir)

	t.Run("with output file", func(t *testing.T) {
		output := path.Join(tmpdir, "test.yml")
		cmd := newCmdRoot()
		cmd.SetArgs([]string{"generate", "source", "test", "--output", output})
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}

		// check the generated config
		cfg, err := ioutil.ReadFile(output)
		if err != nil {
			t.Fatalf("error reading config file output: %v ", err)
		}
		wantConfig := `
# This is an example config file for the test plugin.
account_ids: []
`
		if diff := cmp.Diff(string(cfg), wantConfig); diff != "" {
			t.Errorf("generated config not as expected (+got, -want): %v", diff)
		}
	})
}
