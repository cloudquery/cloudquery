package ec2

import (
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func TransitGatewayMulticastDomains() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_transit_gateway_multicast_domains",
		Description: `https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_TransitGatewayMulticastDomain.html`,
		Resolver:    fetchEc2TransitGatewayMulticastDomains,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Transform:   transformers.TransformWithStruct(&types.TransitGatewayMulticastDomain{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
