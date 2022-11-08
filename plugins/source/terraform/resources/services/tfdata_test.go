package services_test

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources/services"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
	"github.com/rs/zerolog"
)

func TestTfData(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	pth := filepath.Dir(filename)
	p := plugins.NewSourcePlugin(
		"terraform",
		"test",
		[]*schema.Table{
			services.TFData(),
		},
		client.Configure,
	)
	logger := zerolog.New(zerolog.NewTestWriter(t)).Output(
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.StampMicro},
	).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	p.SetLogger(logger)
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:         "dev",
		Version:      "vDev",
		Path:         "test/dev",
		Tables:       []string{"*"},
		Destinations: []string{"mock-destination"},
		Spec: map[string]interface{}{
			"backends": []map[string]interface{}{
				{
					"name": "mylocal",
					"type": "local",
					"config": map[string]string{
						"path": path.Join(pth, "testdata/terraform.tfstate"),
					},
				},
			},
		},
	})
}
