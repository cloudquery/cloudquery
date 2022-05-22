package resources

import (
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestTfData(t *testing.T) {
	cfg := `
	config "mylocal" {
      backend = "local"
      path = "../examples/terraform.tfstate"
    }
`

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: Provider(),
		Config:   cfg,
	})
}
