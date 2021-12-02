package integration_tests

import (
	"testing"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEcsTaskDefinitions(t *testing.T) {
	resource := resources.EcsTaskDefinitions()
	awsTestIntegrationHelper(t, resource, []string{"aws_ecs_clusters.tf", "aws_vpc.tf", "aws_elbv2_load_balancers.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resource.Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"compatibilities":          []interface{}{"EC2"},
						"cpu":                      "1024",
						"family":                   "openapi-task-definition",
						"memory":                   "2048",
						"network_mode":             "awsvpc",
						"requires_compatibilities": []interface{}{"EC2"},
						"revision":                 float64(33),
						"status":                   "ACTIVE",
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_ecs_task_definition_container_definitions",
					ForeignKeyName: "task_definition_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"cpu":           float64(10),
								"docker_labels": nil,
								"essential":     true,
								"image":         "nginx",
							},
						},
					},
				},
				{
					Name:           "aws_ecs_task_definition_requires_attributes",
					ForeignKeyName: "task_definition_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 4,
							Data:  map[string]interface{}{},
						},
					},
				},
			},
		}
	})
}
