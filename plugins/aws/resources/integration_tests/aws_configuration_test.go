package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationConfigConfigurationRecorders(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ConfigConfigurationRecorders(), []string{"aws_configuration.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_config_configuration_recorders",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("config-cr-%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name":                          fmt.Sprintf("config-cr-%s-%s", res.Prefix, res.Suffix),
						"recording_group_all_supported": true,
						"recording_group_include_global_resource_types": false,
					},
				},
			},
		}
	})
}

func TestIntegrationConfigConformancePack(t *testing.T) {
	awsTestIntegrationHelper(t, resources.ConfigConformancePack(), []string{"aws_configuration.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_config_conformance_packs",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("conformance_pack_name = ?", fmt.Sprintf("config-cp-%s-%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"conformance_pack_name":  fmt.Sprintf("config-cp-%s-%s", res.Prefix, res.Suffix),
						"created_by":             nil,
						"delivery_s3_bucket":     nil,
						"delivery_s3_key_prefix": nil,
						"conformance_pack_input_parameters": map[string]interface{}{
							"AccessKeysRotatedParameterMaxAccessKeyAge": "90",
						},
					},
				},
			},
		}
	})
}
