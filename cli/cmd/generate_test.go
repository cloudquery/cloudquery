package cmd

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func TestGenerate(t *testing.T) {
	tmpdir, tmpErr := os.MkdirTemp("", "generate_test_*")
	if tmpErr != nil {
		t.Fatalf("failed to create temporary directory: %v", tmpErr)
	}
	defer os.RemoveAll(tmpdir)

	t.Run("with output file", func(t *testing.T) {
		output := path.Join(tmpdir, "test.yml")
		cmd := NewCmdGenerate()
		cmd.SetArgs([]string{"source", "test", "--output", output})
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}

		// sanity check that valid yaml has been produced
		var cfg interface{}
		yamlFile, err := ioutil.ReadFile(output)
		if err != nil {
			t.Fatalf("error reading config file output: %v ", err)
		}
		err = yaml.Unmarshal(yamlFile, &cfg)
		if err != nil {
			t.Fatalf("error unmarshaling config file: %v", err)
		}
	})
}
