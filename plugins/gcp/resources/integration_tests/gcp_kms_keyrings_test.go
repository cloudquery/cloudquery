package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationKmsKeyrings(t *testing.T) {
	testIntegrationHelper(t, resources.KmsKeyrings(), []string{
		"gcp_kms_keyrings.tf",
		"service-account.tf",
	}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.KmsKeyrings().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Like{"name": fmt.Sprintf("%%kms-keyring-%s%s-v2", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"location": "global",
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "gcp_kms_keyring_crypto_keys",
					ForeignKeyName: "keyring_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"location":                          "global",
								"version_template_algorithm":        "GOOGLE_SYMMETRIC_ENCRYPTION",
								"primary_protection_level":          "SOFTWARE",
								"primary_state":                     "ENABLED",
								"purpose":                           "ENCRYPT_DECRYPT",
								"rotation_period":                   "100000s",
								"primary_algorithm":                 "GOOGLE_SYMMETRIC_ENCRYPTION",
								"version_template_protection_level": "SOFTWARE",
							},
						},
					},
				},
			},
		}
	})
}
