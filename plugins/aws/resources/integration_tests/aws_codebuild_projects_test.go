package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationCodebuildProjects(t *testing.T) {
	resource := resources.CodebuildProjects()
	awsTestIntegrationHelper(t, resource, []string{"aws_codebuild_projects.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resource.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("project-%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"project_visibility":                            "PRIVATE",
						"queued_timeout_in_minutes":                     float64(480),
						"source_type":                                   "GITHUB",
						"source_git_clone_depth":                        float64(1),
						"source_git_submodules_config_fetch_submodules": true,
						"source_insecure_ssl":                           false,
						"source_location":                               "https://github.com/mitchellh/packer.git",
						"source_report_build_status":                    false,
						"source_version":                                "master",
						"tags": map[string]interface{}{
							"Type": "integration_test", "TestId": "windowsfifl5fe", "Environment": "Test",
						},
						"timeout_in_minutes": float64(5),
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_codebuild_project_environment_variables",
					ForeignKeyName: "project_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name":  "SOME_KEY1",
								"value": "SOME_VALUE1",
								"type":  "PLAINTEXT",
							},
						},
					},
				},
				{
					Name:           "aws_codebuild_project_environment_variables",
					ForeignKeyName: "project_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"name":  "SOME_KEY2",
								"value": "SOME_VALUE2",
								"type":  "PARAMETER_STORE",
							},
						},
					},
				},
				{
					Name:           "aws_codebuild_project_secondary_sources",
					ForeignKeyName: "project_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"type":              "S3",
								"insecure_ssl":      false,
								"source_identifier": "package",
							},
						},
					},
				},
				{
					Name:           "aws_codebuild_project_secondary_artifacts",
					ForeignKeyName: "project_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"type":                   "S3",
								"encryption_disabled":    false,
								"artifact_identifier":    "package",
								"namespace_type":         "NONE",
								"override_artifact_name": false,
							},
						},
					},
				},
				{
					Name:           "aws_codebuild_project_file_system_locations",
					ForeignKeyName: "project_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"identifier":  "CODEBUILD_MY_EFS",
								"type":        "EFS",
								"mount_point": "/path",
							},
						},
					},
				},
			},
		}
	})
}
