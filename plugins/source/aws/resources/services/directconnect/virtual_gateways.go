package directconnect

import (
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VirtualGateways() *schema.Table {
	tableName := "aws_directconnect_virtual_gateways"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_VirtualGateway.html`,
		Resolver:    fetchDirectconnectVirtualGateways,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "directconnect"),
		Transform:   client.TransformWithStruct(&types.VirtualGateway{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
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
