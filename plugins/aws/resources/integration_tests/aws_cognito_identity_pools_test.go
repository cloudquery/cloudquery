package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationCognitoIdentityPools(t *testing.T) {
	awsTestIntegrationHelper(t, resources.CognitoIdentityPools(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_cognito_identity_pools",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("identity_pool_name = ?", fmt.Sprintf("cognito_identity_pool%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"identity_pool_name":               fmt.Sprintf("cognito_identity_pool%s-%s", res.Prefix, res.Suffix),
						"allow_unauthenticated_identities": false,
						"allow_classic_flow":               false,
					},
				},
			},
		}
	})
}
