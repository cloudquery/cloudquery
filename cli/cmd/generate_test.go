package cmd

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const wantSourceConfig = `
kind: "source"
spec:
  # Name of the plugin.
  name: "test"

  # Version of the plugin to use.
  version: "v1.1.4"

  # Registry to use (one of "github", "local" or "grpc").
  registry: "github"

  # Path to plugin. Required format depends on the registry.
  path: "cloudquery/test"

  # List of tables to sync.
  tables: ["*"]

  ## Tables to skip during sync. Optional.
  # skip_tables: []

  # Names of destination plugins to sync to.
  destinations: ["postgresql"]

  ## Approximate cap on number of requests to perform concurrently. Optional.
  # concurrency: 5

  # Plugin-specific configuration.
  spec:
    
    # This is an example config file for the test plugin.
    account_ids: []
`

const wantDestinationConfig = `
kind: "destination"
spec:
  # Name of the plugin.
  name: "postgresql"

  # Version of the plugin to use.
  version: "v0.0.1"

  # Registry to use (one of "github", "local" or "grpc").
  registry: "github"

  # Path to plugin. Required format depends on the registry.
  path: "cloudquery/postgresql"

  # Write mode (either "overwrite" or "append").
  write_mode: "append"

  # Plugin-specific configuration.
  spec:
    
    connection_string: "postgresql://user:password@localhost:5432/dbname"
`

func TestGenerate(t *testing.T) {
	tmpdir, tmpErr := os.MkdirTemp("", "generate_test_*")
	if tmpErr != nil {
		t.Fatalf("failed to create temporary directory: %v", tmpErr)
	}
	defer os.RemoveAll(tmpdir)

	t.Run("generate source", func(t *testing.T) {
		output := path.Join(tmpdir, "test-source.yml")
		cmd := NewCmdRoot()

		cmd.SetArgs([]string{"generate", "source", "test", "--output", output})
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}

		// check the generated config
		cfg, err := os.ReadFile(output)
		if err != nil {
			t.Fatalf("error reading config file output: %v ", err)
		}
		cfgStr := strings.TrimSpace(string(cfg))
		want := strings.TrimSpace(wantSourceConfig)
		if diff := cmp.Diff(cfgStr, want); diff != "" {
			t.Errorf("generated source config not as expected (-got, +want): %v", diff)
		}
	})

	t.Run("generate destination", func(t *testing.T) {
		// TODO: Change this to use a test destination plugin when we have one.
		//       For now, this test can be manually run against the postgresql
		//       plugin by commenting out this skip line.
		t.Skip()

		output := path.Join(tmpdir, "test-destination.yml")
		cmd := NewCmdRoot()
		cmd.SetArgs([]string{"generate", "destination", "postgresql", "--output", output})
		if err := cmd.Execute(); err != nil {
			t.Fatal(err)
		}

		// check the generated config
		cfg, err := os.ReadFile(output)
		if err != nil {
			t.Fatalf("error reading config file output: %v ", err)
		}
		cfgStr := strings.TrimSpace(string(cfg))
		want := strings.TrimSpace(wantDestinationConfig)
		if diff := cmp.Diff(cfgStr, want); diff != "" {
			t.Errorf("generated source config not as expected (-got, +want): %v", diff)
		}
	})
}
