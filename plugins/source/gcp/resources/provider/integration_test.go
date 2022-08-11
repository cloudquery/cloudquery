//go:build integration

package provider

import (
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegration(t *testing.T) {
	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider: Provider(),
	})
}
