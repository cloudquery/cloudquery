package lambda

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Layers() *schema.Table {
	tableName := "aws_lambda_layers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_LayersListItem.html`,
		Resolver:    fetchLambdaLayers,
		Transform:   transformers.TransformWithStruct(&types.LayersListItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lambda"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LayerArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			layerVersions(),
		},
	}
}
