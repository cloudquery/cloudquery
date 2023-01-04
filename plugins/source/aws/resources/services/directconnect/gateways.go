package directconnect

import (
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Gateways() *schema.Table {
	return &schema.Table{
		Name:        "aws_directconnect_gateways",
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGateway.html`,
		Resolver:    fetchDirectconnectGateways,
		Multiplex:   client.ServiceAccountRegionMultiplexer("directconnect"),
		Transform:  transformers.TransformWithStruct(&types.DirectConnectGateway{}),
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
				Resolver: resolveGatewayARN(),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DirectConnectGatewayId"),
			},
		},
		Relations: []*schema.Table{
			GatewayAssociations(),
			GatewayAttachments(),
		},
	}
}
