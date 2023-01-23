package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TransitGatewayRouteTables() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_transit_gateway_route_tables",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayRouteTable.html`,
		Resolver:    fetchEc2TransitGatewayRouteTables,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.TransitGatewayRouteTable{}),
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
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}
