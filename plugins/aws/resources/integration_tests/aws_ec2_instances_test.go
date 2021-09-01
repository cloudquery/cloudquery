package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2Instances(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Instances(), []string{"aws_ec2_instances.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2Instances().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{
						"tags->>'TestId'": res.Suffix,
						"tags->>'Name'":   fmt.Sprintf("ec2_instance%s", res.Suffix),
					},
					squirrel.NotEq{"state_name": "terminated"},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"virtualization_type":          "hvm",
						"instance_type":                "t2.nano",
						"root_device_type":             "ebs",
						"source_dest_check":            false,
						"ena_support":                  true,
						"cpu_options_threads_per_core": float64(1),
						"cpu_options_core_count":       float64(1),
						"architecture":                 "x86_64",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("ec2_instance%s", res.Suffix),
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_ec2_instance_block_device_mappings",
					ForeignKeyName: "instance_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"ebs_status":                "attached",
							"ebs_delete_on_termination": true,
						},
					}},
				},
				{
					Name:           "aws_ec2_instance_network_interfaces",
					ForeignKeyName: "instance_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"attachment_status":                "attached",
							"interface_type":                   "interface",
							"attachment_delete_on_termination": true,
							"source_dest_check":                false,
						},
					}},
					Relations: []*providertest.ResourceIntegrationVerification{
						{
							Name:           "aws_ec2_instance_network_interface_groups",
							ForeignKeyName: "instance_network_interface_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"group_name": fmt.Sprintf("aws_ec2_instances_sg_%s%s", res.Prefix, res.Suffix),
								},
							}},
						},
						{
							Name:           "aws_ec2_instance_network_interface_private_ip_addresses",
							ForeignKeyName: "instance_network_interface_cq_id",
							ExpectedValues: []providertest.ExpectedValue{{
								Count: 1,
								Data: map[string]interface{}{
									"is_primary": true,
								},
							}},
						},
					},
				},
				{
					Name:           "aws_ec2_instance_security_groups",
					ForeignKeyName: "instance_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"group_name": fmt.Sprintf("aws_ec2_instances_sg_%s%s", res.Prefix, res.Suffix),
						},
					}},
				},
			},
		}
	})
}
