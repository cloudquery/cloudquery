package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationDirectconnectLags(t *testing.T) {
	awsTestIntegrationHelper(t, resources.DirectconnectLags(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_directconnect_lags",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{squirrel.Eq{"name": fmt.Sprintf("dx-lag-%s-%s", res.Prefix, res.Suffix)}, squirrel.NotEq{"state": "deleted"}})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                  fmt.Sprintf("dx-lag-%s-%s", res.Prefix, res.Suffix),
						"connections_bandwidth": "1Gbps",
						"jumbo_frame_capable":   true,
						"location":              "EqDC2",
						"mac_sec_capable":       false,
					},
				},
			},
		}
	})
}
