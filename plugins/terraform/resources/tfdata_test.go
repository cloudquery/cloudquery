package resources

import (
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
	"github.com/cloudquery/cq-provider-terraform/client"
)

type TestLocalConfigBlock struct {
	Config []TestLocalConfig `hcl:"config,block"`
}

type TestLocalConfig struct {
	BackendType string `hcl:"backend,attr"`
	BackendName string `hcl:"config,label"`
	Path        string `hcl:"path,attr"`
}

func TestTfData(t *testing.T) {
	cfg := TestLocalConfigBlock{
		Config: []TestLocalConfig{
			{
				BackendType: "local",
				BackendName: "mylocal",
				Path:        "../examples/terraform.tfstate",
			},
		},
	}

	providertest.TestResource(t, Provider, providertest.ResourceTestData{
		Table:     TFData(),
		Config:    cfg,
		Configure: client.Configure,
	})

}
