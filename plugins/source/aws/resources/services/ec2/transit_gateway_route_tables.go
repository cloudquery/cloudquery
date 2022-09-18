// Code generated by codegen; DO NOT EDIT.

package ec2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func TransitGatewayRouteTables() *schema.Table {
	return &schema.Table{
		Name:      "aws_ec2_transit_gateway_route_tables",
		Resolver:  fetchEc2TransitGatewayRouteTables,
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
				Name:     "transit_gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
			{
				Name:     "creation_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreationTime"),
			},
			{
				Name:     "default_association_route_table",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultAssociationRouteTable"),
			},
			{
				Name:     "default_propagation_route_table",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultPropagationRouteTable"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "transit_gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayId"),
			},
			{
				Name:     "transit_gateway_route_table_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TransitGatewayRouteTableId"),
			},
		},
	}
}
