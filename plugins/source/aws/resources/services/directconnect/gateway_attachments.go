package directconnect

import (
	"github.com/aws/aws-sdk-go-v2/service/directconnect/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GatewayAttachments() *schema.Table {
	tableName := "aws_directconnect_gateway_attachments"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/directconnect/latest/APIReference/API_DirectConnectGatewayAttachment.html`,
		Resolver:    fetchDirectconnectGatewayAttachments,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "directconnect"),
		Transform:   transformers.TransformWithStruct(&types.DirectConnectGatewayAttachment{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
