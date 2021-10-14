package integration_tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/cloudquery/cq-provider-azure/client"
	"github.com/cloudquery/cq-provider-azure/resources"

	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

// IntegrationTestsEnabledVar is the name of the environment variable that enables integration tests from this package.
// Set it to one of "1", "y", "yes", "true" to enable the tests.
const IntegrationTestsEnabledVar = "INTEGRATION_TESTS"

func lastN(str string, i int) string {
	if len(str) <= i {
		return str
	}

	return str[len(str)-i:]
}

func awsTestIntegrationHelper(t *testing.T, table *schema.Table, resourceFiles []string, verificationBuilder func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification) {
	cfg := client.Config{}

	providertest.IntegrationTest(t, resources.Provider, providertest.ResourceIntegrationTestData{
		Table:               table,
		Config:              cfg,
		Configure:           client.Configure,
		VerificationBuilder: verificationBuilder,
		Resources:           resourceFiles,
	})
}

func TestMain(m *testing.M) {
	enabled := os.Getenv(IntegrationTestsEnabledVar)
	enabledValues := map[string]struct{}{
		"1":       {},
		"y":       {},
		"yes":     {},
		"true":    {},
		"enable":  {},
		"enabled": {},
	}
	if _, ok := enabledValues[enabled]; ok {
		os.Exit(m.Run())
	} else {
		fmt.Fprintln(os.Stderr, "Integration tests are skipped. Set INTEGRATION_TESTS=1 environment variable to enable.")
	}
}
