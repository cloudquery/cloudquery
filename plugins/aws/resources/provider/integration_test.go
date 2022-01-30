//go:build integration

package provider

import (
	"os"
	"testing"

	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegration(t *testing.T) {
	cfg := `
		max_retries = 20
  		max_backoff = 60
  	`
	// we want to give a way for the user ci to override with additional cfg for integration test.
	// For example:
	//
	// accounts "cq-provider-aws" {
	// 	role_arn = "arn:aws:iam::70xxxxxxxxxx:role/CqProviderxxxxxxxxxxxxxxxxx"
	// }
	// accounts "cq-dev" {
	//  role_arn = "arn:aws:iam::70xxxxxxxxxx:role/CqProviderxxxxxxxxxxxxxxxxx"
	// }

	additionalConfig := os.Getenv("CQ_TEST_CFG")
	if additionalConfig != "" {
		cfg += additionalConfig
	}

	providertest.TestResource(t, providertest.ResourceTestCase{
		Provider:              Provider(),
		Config:                cfg,
		NotParallel:           true,
		ParallelFetchingLimit: 10000,
	})
}
