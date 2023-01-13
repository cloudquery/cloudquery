package directconnect

import (
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VirtualGateways() *schema.Table {
	return &schema.Table{
		Name:        "aws_directconnect_virtual_gateways",
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualGateway.html`,
		Resolver:    fetchDirectconnectVirtualGateways,
		Multiplex:   client.ServiceAccountRegionMultiplexer("directconnect"),
		Transform:   transformers.TransformWithStruct(&types.VirtualGateway{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VirtualGatewayId"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
