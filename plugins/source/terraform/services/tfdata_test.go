package services_test

import (
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/terraform/client"
	"github.com/cloudquery/cloudquery/plugins/source/terraform/services"
	"github.com/cloudquery/plugin-sdk/plugins"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/specs"
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
	plugins.TestSourcePluginSync(t, p, specs.Source{
		Name:         "dev",
		Tables:       []string{"*"},
		Destinations: []string{},
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
