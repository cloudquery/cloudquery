package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationDirectconnectConnections(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DirectconnectConnections(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_directconnect_connections",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{squirrel.Eq{"name": fmt.Sprintf("dx-connection%s-%s", res.Prefix, res.Suffix)}, squirrel.NotEq{"connection_state": "deleted"}})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                fmt.Sprintf("dx-connection%s-%s", res.Prefix, res.Suffix),
						"bandwidth":           "1Gbps",
						"jumbo_frame_capable": false,
						"location":            "EqDC2",
						"mac_sec_capable":     false,
					},
				},
			},
		}
	})
}
