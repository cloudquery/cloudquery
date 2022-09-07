package resources_test

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/terraform/resources/provider"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestTfData(t *testing.T) {
	cfg := `
    config:
      - name: mylocal
        backend: local
        path: "../examples/terraform.tfstate"
`

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: provider.Provider(),
		Config:   cfg,
	})
}
