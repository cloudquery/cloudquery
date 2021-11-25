package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationAutoscalingGroups(t *testing.T) {
	resource := resources.AutoscalingGroups()
	awsTestIntegrationHelper(t, resource, []string{"aws_autoscaling_groups.tf", "aws_vpc.tf", "aws_elbv1_load_balancers.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resource.Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where("name = ?", fmt.Sprintf("ag-%s%s", res.Prefix, res.Suffix))
			},
			ExpectedValues: []providertest.ExpectedValue{{
				Count: 1,
				Data: map[string]interface{}{
					"availability_zones":                    []interface{}{"us-east-1a", "us-east-1e"},
					"default_cooldown":                      float64(300),
					"health_check_grace_period":             float64(300),
					"desired_capacity":                      float64(1),
					"max_size":                              float64(1),
					"min_size":                              float64(1),
					"capacity_rebalance":                    true,
					"new_instances_protected_from_scale_in": false,
					"health_check_type":                     "ELB",
					"termination_policies":                  []interface{}{"Default"},
				},
			}},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_autoscaling_group_instances",
					ForeignKeyName: "group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"protected_from_scale_in": false,
							"type":                    "t2.nano",
						},
					}},
				},
				{
					Name:           "aws_autoscaling_group_lifecycle_hooks",
					ForeignKeyName: "group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"default_result":        "CONTINUE",
							"global_timeout":        float64(172800),
							"heartbeat_timeout":     float64(2000),
							"lifecycle_hook_name":   fmt.Sprintf("foobar%s%s", res.Prefix, res.Suffix),
							"lifecycle_transition":  "autoscaling:EC2_INSTANCE_LAUNCHING",
							"notification_metadata": "{\"foo\":\"bar\"}",
						},
					}},
				},
				{
					Name:           "aws_autoscaling_group_tags",
					ForeignKeyName: "group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"key":                 "foo",
							"propagate_at_launch": true,
							"resource_type":       "auto-scaling-group",
							"value":               "bar",
						},
					}},
				},
				{
					Name:           "aws_autoscaling_group_tags",
					ForeignKeyName: "group_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"key":                 "lorem",
							"propagate_at_launch": false,
							"resource_type":       "auto-scaling-group",
							"value":               "ipsum",
						},
					}},
				},
			},
		}
	})
}
