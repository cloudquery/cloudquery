// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func EgressOnlyInternetGateways() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_egress_only_internet_gateways",
		Resolver:  fetchEc2EgressOnlyInternetGateways,
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
				Resolver: resolveEgressOnlyInternetGatewaysArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "attachments",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attachments"),
			},
			{
				Name:     "egress_only_internet_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EgressOnlyInternetGatewayId"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}
