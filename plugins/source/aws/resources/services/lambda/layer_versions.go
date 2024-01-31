package lambda

import (
	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func layerVersions() *schema.Table {
	tableName := "aws_lambda_layer_versions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_LayerVersionsListItem.html`,
		Resolver:    fetchLambdaLayerVersions,
		Transform:   transformers.TransformWithStruct(&types.LayerVersionsListItem{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("LayerVersionArn"),
				PrimaryKeyComponent: true,
			},
			{
				Name:     "layer_arn",
				Type:     arrow.BinaryTypes.String,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},

		Relations: []*schema.Table{
			layerVersionPolicies(),
		},
	}
}
