package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-aws/resources"
	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func awsTestIntegrationHelper(t *testing.T, table *schema.Table, verificationBuilder func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification) {
	cfg := client.Config{
		Regions:  []string{"us-east-1"},
		AWSDebug: false,
	}

	providertest.IntegrationTest(t, resources.Provider, providertest.ResourceIntegrationTestData{
		Table:               table,
		Config:              cfg,
		Configure:           client.Configure,
		VerificationBuilder: verificationBuilder,
	})
}
