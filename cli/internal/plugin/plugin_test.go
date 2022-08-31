package plugin

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

func TestPluginManagerDownloadSource(t *testing.T) {
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	dirName, err := ioutil.TempDir(os.TempDir(), "cq-plugins")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dirName)

	pm := NewPluginManager(WithDirectory(dirName), WithLogger(l))
	if _, err := pm.DownloadSource(ctx, specs.Source{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "cloudquery/test",
		Version:  "latest",
	}); err != nil {
		t.Fatal(fmt.Errorf("failed to download official source plugin test latest: %w", err))
	}

	if _, err := pm.DownloadSource(ctx, specs.Source{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "cloudquery/test",
		Version:  "v1.1.0",
	}); err != nil {
		t.Fatal(fmt.Errorf("failed to download official source plugin test v1.1.0: %w", err))
	}

	if _, err := pm.DownloadSource(ctx, specs.Source{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "yevgenypats/test",
		Version:  "v1.0.0",
	}); err != nil {
		t.Fatal(fmt.Errorf("failed to download community source plugin test v1.0.0: %w", err))
	}

	if _, err := pm.DownloadSource(ctx, specs.Source{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "yevgenypats/test",
		Version:  "latest",
	}); err != nil {
		t.Fatal(fmt.Errorf("failed to download community source plugin test latest %w", err))
	}

}

func TestPluginManagerGetSourceClient(t *testing.T) {
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	dirName, err := ioutil.TempDir(os.TempDir(), "cq-plugins")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dirName)
	pm := NewPluginManager(WithDirectory(dirName), WithLogger(l))
	pl, err := pm.NewSourcePlugin(ctx, specs.Source{
		Name:     "test",
		Registry: specs.RegistryGithub,
		Path:     "cloudquery/test",
		Version:  "latest",
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

}
