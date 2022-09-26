package plugins

import (
	"context"
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
		Version:  "v1.1.5",
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


func TestPluginManagerGetSourceClient(t *testing.T) {
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	for _, tc := range getSourceClientTestCases {
		t.Run(tc.Path + "_" + tc.Version, func(t *testing.T) {
			dirName := t.TempDir()
			pm := NewPluginManager(WithDirectory(dirName), WithLogger(l))
			pl, err := pm.NewSourcePlugin(ctx, tc.Registry, tc.Path, tc.Version)
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


var getDestinationClientTestCases = []specs.Destination{
	{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "cloudquery/test",
		Version:  "latest",
	},
	// {
	// 	Name:     "test",
	// 	Registry: specs.RegistryGithub,
	// 	Path:     "cloudquery/test",
	// 	Version:  "v1.1.0",
	// },
	// {
	// 	Name:     "test",
	// 	Registry: specs.RegistryGithub,
	// 	Path:     "yevgenypats/test",
	// 	Version:  "latest",
	// },
	// {
	// 	Name:     "test",
	// 	Registry: specs.RegistryGithub,
	// 	Path:     "yevgenypats/test",
	// 	Version:  "v1.0.1",
	// },
}

func TestPluginManagerGetDestinationClient(t *testing.T) {
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	dirName := t.TempDir()
	pm := NewPluginManager(WithDirectory(dirName), WithLogger(l))
	for _, tc := range getDestinationClientTestCases {
		t.Run(tc.Path + "_" + tc.Version, func(t *testing.T) {
			pl, err := pm.NewDestinationPlugin(ctx, tc.Registry, tc.Path, tc.Version)
			if err != nil {
				t.Fatal(err)
			}
			defer pl.Close()
			client := pl.GetClient()
			name, err := client.Name(ctx)
			if err != nil {
				t.Fatal(err)
			}
			if name != "test" {
				t.Fatal("expected name to be test got ", name)
			}
		})
	}
}