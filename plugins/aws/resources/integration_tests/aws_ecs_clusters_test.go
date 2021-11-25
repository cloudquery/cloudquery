package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEcsClusters(t *testing.T) {
	awsTestIntegrationHelper(t, resources.EcsClusters(), []string{"aws_ecs_clusters.tf", "aws_vpc.tf", "aws_elbv2_load_balancers.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ecs_clusters",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.Eq{"name": fmt.Sprintf("ecs_cluster_%s%s", res.Prefix, res.Suffix)})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"name": fmt.Sprintf("ecs_cluster_%s%s", res.Prefix, res.Suffix),
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"name":   fmt.Sprintf("ecs_cluster_%s%s", res.Prefix, res.Suffix),
							"stage":  "test",
							"TestId": res.Suffix,
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_ecs_cluster_services",
					ForeignKeyName: "cluster_cq_id",
					ExpectedValues: []providertest.ExpectedValue{
						{
							Count: 1,
							Data: map[string]interface{}{
								"desired_count":                            float64(1),
								"enable_ecs_managed_tags":                  false,
								"launch_type":                              "EC2",
								"deployment_configuration_maximum_percent": float64(200),
								"deployment_configuration_deployment_circuit_breaker_rollback": false,
								"deployment_configuration_deployment_circuit_breaker_enable":   false,
								"enable_execute_command":            false,
								"health_check_grace_period_seconds": float64(0),
							},
						},
					},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "aws_ecs_cluster_service_load_balancers",
							ForeignKeyName: "cluster_service_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"container_name": "web-server",
										"container_port": float64(8080),
									},
								},
							},
						},
						{
							Name:           "aws_ecs_cluster_service_deployments",
							ForeignKeyName: "cluster_service_cq_id",
							ExpectedValues: []providertest.ExpectedValue{
								{
									Count: 1,
									Data: map[string]interface{}{
										"launch_type": "EC2",
									},
								},
							},
						},
					},
				},
			},
		}
	})
}
