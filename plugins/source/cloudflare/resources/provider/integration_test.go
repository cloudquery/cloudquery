//go:build integration

package provider

import (
	"os"
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegration(t *testing.T) {
	cfg := `
  	`

	additionalConfig := os.Getenv("CQ_TEST_CFG")
	if additionalConfig != "" {
		cfg += additionalConfig
	}

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider:              Provider(),
		Config:                cfg,
		NotParallel:           true,
		ParallelFetchingLimit: 10000,
		SkipIgnoreInTest:      false,
	})
}
