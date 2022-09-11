// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func NetworkAcls() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_network_acls",
		Resolver:  fetchEc2NetworkAcls,
		Multiplex: client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveNetworkAclArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "connectivity_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConnectivityType"),
			},
			{
				Name:     "create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreateTime"),
			},
			{
				Name:     "delete_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("DeleteTime"),
			},
			{
				Name:     "failure_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureCode"),
			},
			{
				Name:     "failure_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("FailureMessage"),
			},
			{
				Name:     "nat_gateway_addresses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NatGatewayAddresses"),
			},
			{
				Name:     "nat_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("NatGatewayId"),
			},
			{
				Name:     "provisioned_bandwidth",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ProvisionedBandwidth"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "subnet_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SubnetId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}
