package integration_tests

import (
	"fmt"
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-sdk/provider/providertest"
)

func TestIntegrationDirectConnectGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DirectconnectGateways(), func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_directconnect_gateways",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("direct_connect_gateway_name = ?", fmt.Sprintf("%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"amazon_side_asn": float64(64512),
				},
			}},
			//Relations: []*providertest.ResourceIntegrationVerification{
			//	{
			//		Name:           "aws_lambda_function_aliases",
			//		ForeignKeyName: "function_id",
			//		ExpectedValues: []providertest.ExpectedValue{{
			//			Count: 1,
			//			Data: map[string]interface{}{
			//				"description": "a sample description",
			//			},
			//		}},
			//	},
			//},
		}
	})
}
