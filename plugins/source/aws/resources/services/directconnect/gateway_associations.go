package directconnect

import (
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GatewayAssociations() *schema.Table {
	return &schema.Table{
		Name:        "aws_directconnect_gateway_associations",
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAssociation.html`,
		Resolver:    fetchDirectconnectGatewayAssociations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("directconnect"),
		Transform:   transformers.TransformWithStruct(&types.DirectConnectGatewayAssociation{}),
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
				Name:     "gateway_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "gateway_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
		},
	}
}
