package appstream

import (
	"github.com/aws/aws-sdk-go-v2/service/appstream/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func AppBlocks() *schema.Table {
	tableName := "aws_appstream_app_blocks"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/appstream2/latest/APIReference/API_AppBlock.html`,
		Resolver:    fetchAppstreamAppBlocks,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "appstream2"),
		Transform:   client.TransformWithStruct(&types.AppBlock{}),
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
		},
	}
}
