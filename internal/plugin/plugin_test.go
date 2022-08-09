package plugin

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/cloudquery/plugin-sdk/spec"
	"github.com/rs/zerolog"
)

func TestPluginManagerDownload(t *testing.T) {
	ctx := context.Background()
	l := zerolog.New(zerolog.NewTestWriter(t)).Output(zerolog.ConsoleWriter{Out: os.Stderr}).Level(zerolog.DebugLevel)
	dirName, err := ioutil.TempDir(os.TempDir(), "cq-plugins")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dirName)

	pm := NewPluginManager(WithDirectory(dirName), WithLogger(l))
	if err := pm.Download(ctx, spec.SourceSpec{
		Name:     "aws",
		Registry: "github",
		Path:     "cloudquery/aws",
		Version:  "v0.13.6",
	}); err != nil {
		t.Fatal(err)
	}
}
