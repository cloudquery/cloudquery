package plugin

import (
	"context"
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
	if _, err := pm.DownloadSource(ctx, specs.SourceSpec{
		Name:     "test",
		Registry: "github",
		Path:     "cloudquery/test",
		Version:  "v0.0.4",
	}); err != nil {
		t.Fatal(err)
	}
}

func TestPluginManagerGetSourceClient(t *testing.T) {
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	dirName, err := ioutil.TempDir(os.TempDir(), "cq-plugins")
	if err != nil {
		t.Fatal(err)
	}
	// defer os.RemoveAll(dirName)

	pm := NewPluginManager(WithDirectory(dirName), WithLogger(l))
	_, err = pm.GetSourcePluginClient(ctx, specs.SourceSpec{
		Name:     "test",
		Registry: "github",
		Path:     "cloudquery/test",
		Version:  "v0.0.4",
	})
	if err != nil {
		t.Fatal(err)
	}
	defer pm.CloseAll(ctx)
}
