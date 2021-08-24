package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationIamServerCertificates(t *testing.T) {
	awsTestIntegrationHelper(t, resources.IamServerCertificates(), nil, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_iam_server_certificates",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("iam_server_certificates_%s%s", res.Prefix, res.Suffix),
						"path": "/",
					},
				},
			},
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("iam_server_certificates_%s%s", res.Prefix, res.Suffix)})
			},
		}
	})
}
