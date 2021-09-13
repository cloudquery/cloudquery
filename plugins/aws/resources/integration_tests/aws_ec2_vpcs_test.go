package integration_tests

import (
	"fmt"
	"testing"

	"github.com/Masterminds/squirrel"

	"github.com/cloudquery/cq-provider-aws/resources"
	providertest "github.com/cloudquery/cq-provider-sdk/provider/testing"
)

func TestIntegrationEc2Vpcs(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Vpcs(), []string{"aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_vpcs",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"is_default": false,
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("vpc%s-%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}

func TestIntegrationEc2VpcPeeringConnections(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcPeeringConnections(), []string{"aws_ec2_vpc.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_vpc_peering_connections",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"tags->>'TestId'": res.Suffix},
					squirrel.NotEq{"status_code": "deleted"},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"accepter_allow_dns_resolution_from_remote_vpc":           true,
						"accepter_allow_egress_local_classic_link_to_remote_vpc":  false,
						"accepter_allow_egress_local_vpc_to_remote_classic_link":  false,
						"requester_allow_dns_resolution_from_remote_vpc":          true,
						"requester_allow_egress_local_classic_link_to_remote_vpc": false,
						"requester_allow_egress_local_vpc_to_remote_classic_link": false,
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}

func TestIntegrationEc2VpcEndpoints(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2VpcEndpoints(), []string{"aws_ec2_vpc.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_vpc_endpoints",
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"vpc_endpoint_type":   "Gateway",
						"requester_managed":   false,
						"private_dns_enabled": false,
						"tags": map[string]interface{}{
							"Type":        "integration_test",
							"Environment": "test",
							"TestId":      res.Suffix,
						},
					},
				},
			},
		}
	})
}

func TestIntegrationEc2TransitGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2TransitGateways(), []string{"aws_ec2_vpc.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_transit_gateways",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"tags->>'TestId'": res.Suffix},
					squirrel.NotEq{"state": "deleted"},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"multicast_support":               "disable",
						"dns_support":                     "enable",
						"default_route_table_propagation": "enable",
						"default_route_table_association": "enable",
						"vpn_ecmp_support":                "enable",
						"auto_accept_shared_attachments":  "disable",
						"description":                     fmt.Sprintf("description %s-%s", res.Prefix, res.Suffix),
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}

func TestIntegrationEc2Subnets(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2Subnets(), []string{"aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: "aws_ec2_subnets",
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"tags->>'TestId'": res.Suffix},
					squirrel.NotEq{"state": "deleted"},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"default_for_az":                  false,
						"map_customer_owned_ip_on_launch": false,
						"map_public_ip_on_launch":         false,
						"assign_ipv6_address_on_creation": false,
						"cidr_block":                      "10.0.1.0/24",
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("vpc-subnet%s-%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
		}
	})
}

func TestIntegrationEc2SecurityGroups(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2SecurityGroups(), []string{"aws_ec2_vpc.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2SecurityGroups().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"group_name": fmt.Sprintf("ec2-sg-%s%s", res.Prefix, res.Suffix),
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}

func TestIntegrationEc2RouteTables(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2RouteTables(), []string{"aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2RouteTables().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
							"Name":   fmt.Sprintf("vpc-routetable%s-%s", res.Prefix, res.Suffix),
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_ec2_route_table_associations",
					ForeignKeyName: "route_table_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 2,
						Data: map[string]interface{}{
							"association_state": "associated",
							"main":              false,
						},
					}},
				},
				{
					Name:           "aws_ec2_route_table_routes",
					ForeignKeyName: "route_table_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"destination_cidr_block": "10.0.0.0/16",
						},
					}},
				},
				{
					Name:           "aws_ec2_route_table_routes",
					ForeignKeyName: "route_table_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"destination_cidr_block": "0.0.0.0/0",
						},
					}},
				},
			},
		}
	})
}

func TestIntegrationEc2NetworkAcls(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2NetworkAcls(), []string{"aws_ec2_vpc.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2NetworkAcls().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"is_default": false,
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"Name":   fmt.Sprintf("ec2-acl-%s-%s", res.Prefix, res.Suffix),
							"TestId": res.Suffix,
						},
					},
				},
			},
			Relations: []*providertest.ResourceIntegrationVerification{
				{
					Name:           "aws_ec2_network_acl_entries",
					ForeignKeyName: "network_acl_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"egress":          true,
							"ipv6_cidr_block": nil,
							"port_range_from": float64(443),
							"port_range_to":   float64(443),
							"protocol":        "6",
							"rule_action":     "allow",
						},
					}},
				},
				{
					Name:           "aws_ec2_network_acl_entries",
					ForeignKeyName: "network_acl_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"egress":          true,
							"ipv6_cidr_block": nil,
							"protocol":        "-1",
							"rule_action":     "deny",
						},
					}},
				},
				{
					Name:           "aws_ec2_network_acl_entries",
					ForeignKeyName: "network_acl_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"egress":          false,
							"ipv6_cidr_block": nil,
							"port_range_from": float64(80),
							"port_range_to":   float64(80),
							"protocol":        "6",
							"rule_action":     "allow",
						},
					}},
				},
				{
					Name:           "aws_ec2_network_acl_entries",
					ForeignKeyName: "network_acl_cq_id",
					ExpectedValues: []providertest.ExpectedValue{{
						Count: 1,
						Data: map[string]interface{}{
							"egress":          false,
							"ipv6_cidr_block": nil,
							"protocol":        "-1",
							"rule_action":     "deny",
						},
					}},
				},
			},
		}
	})
}

func TestIntegrationEc2NatGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2NatGateways(), []string{"aws_ec2_vpc.tf", "aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2NatGateways().Name,
			Filter: func(sq squirrel.SelectBuilder, res *providertest.ResourceIntegrationTestData) squirrel.SelectBuilder {
				return sq.Where(squirrel.And{
					squirrel.Eq{"tags->>'TestId'": res.Suffix},
					squirrel.NotEq{"state": "deleted"},
				})
			},
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"Name":   fmt.Sprintf("ec2-nat-%s-%s", res.Prefix, res.Suffix),
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}

func TestIntegrationEc2InternetGateways(t *testing.T) {
	awsTestIntegrationHelper(t, resources.Ec2InternetGateways(), []string{"aws_vpc.tf"}, func(res *providertest.ResourceIntegrationTestData) providertest.ResourceIntegrationVerification {
		return providertest.ResourceIntegrationVerification{
			Name: resources.Ec2InternetGateways().Name,
			ExpectedValues: []providertest.ExpectedValue{
				{
					Count: 1,
					Data: map[string]interface{}{
						"tags": map[string]interface{}{
							"Type":   "integration_test",
							"TestId": res.Suffix,
						},
					},
				},
			},
		}
	})
}
