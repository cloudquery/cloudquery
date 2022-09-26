package plugins

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

var getSourceClientTestCases = []specs.Source{
	{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "cloudquery/test",
		Version:  "latest",
	},
	{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "cloudquery/test",
		Version:  "v1.1.0",
	},
	{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "yevgenypats/test",
		Version:  "latest",
	},
	{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "yevgenypats/test",
		Version:  "v1.0.1",
	},
}

func TestPluginManagerDownloadSource(t *testing.T) {
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	dirName, err := os.MkdirTemp(os.TempDir(), "cq-plugins")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dirName)

	// this test is mainly a smoke test as we test all permutations in GetSourceClientTest
	// which calls to download anyways.
	pm := NewPluginManager(WithDirectory(dirName), WithLogger(l))
	if _, err := pm.DownloadSource(ctx, &specs.Source{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "cloudquery/test",
		Version:  "latest",
	}); err != nil {
		t.Fatal(fmt.Errorf("failed to download official source plugin test latest: %w", err))
	}
}

func TestPluginManagerGetSourceClient(t *testing.T) {
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	for _, tc := range getSourceClientTestCases {
		t.Run(fmt.Sprintf("%s@%s", tc.Path, tc.Version), func(t *testing.T) {
			dirName, err := os.MkdirTemp(os.TempDir(), "cq-plugins")
			if err != nil {
				t.Fatal(err)
			}
			defer os.RemoveAll(dirName)
			pm := NewPluginManager(WithDirectory(dirName), WithLogger(l))
			pl, err := pm.NewSourcePlugin(ctx, &specs.Source{
				Name:     tc.Name,
				Registry: tc.Registry,
				Path:     tc.Path,
				Version:  tc.Version,
			})
			if err != nil {
				t.Fatal(err)
			}
			defer pl.Close()
			client := pl.GetClient()
			tables, err := client.GetTables(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if len(tables) != 1 {
				t.Fatal("expected 1 table got ", len(tables))
			}
		})
	}
}
