package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-gcp/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationComputeSSLPolicies(t *testing.T) {
	table := resources.ComputeSslPolicies()
	testIntegrationHelper(t, table, nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: table.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":            fmt.Sprintf("prod-ssl-policy-%s%s", res.Prefix, res.Suffix),
						"profile":         "MODERN",
						"kind":            "compute#sslPolicy",
						"custom_features": nil,
						"min_tls_version": "TLS_1_0",
					},
				},
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":            fmt.Sprintf("nonprod-ssl-policy-%s%s", res.Prefix, res.Suffix),
						"profile":         "MODERN",
						"kind":            "compute#sslPolicy",
						"custom_features": nil,
						"min_tls_version": "TLS_1_2",
					},
				},
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":    fmt.Sprintf("custom-ssl-policy-%s%s", res.Prefix, res.Suffix),
						"profile": "CUSTOM",
						"kind":    "compute#sslPolicy",
						"custom_features": []interface{}{
							"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
							"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
						},
						"min_tls_version": "TLS_1_2",
					},
				},
			},
		}
	})
}
