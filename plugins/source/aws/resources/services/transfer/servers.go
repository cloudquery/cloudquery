package transfer

import (
	"github.com/aws/aws-sdk-go-v2/service/transfer/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Servers() *schema.Table {
	return &schema.Table{
		Name:                "aws_transfer_servers",
		Description:         `https://docs.aws.amazon.com/transfer/latest/userguide/API_DescribedServer.html`,
		Resolver:            fetchTransferServers,
		PreResourceResolver: getServer,
		Transform:           transformers.TransformWithStruct(&types.DescribedServer{}),
		Multiplex:           client.ServiceAccountRegionMultiplexer("transfer"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:        "tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveServersTags,
				Description: `Specifies the key-value pairs that you can use to search for and group servers that were assigned to the server that was described`,
			},
		},
	}
}
