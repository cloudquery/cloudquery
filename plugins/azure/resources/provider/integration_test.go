//go:build integration

package provider

import (
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegration(t *testing.T) {
	config := `
		subscriptions = ["78f26f10-0e60-4293-8a7e-122584ccb40d"]
	`
	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: Provider(),
		Config:   config,
	})
}
