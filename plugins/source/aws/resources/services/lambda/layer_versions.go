package lambda

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func LayerVersions() *schema.Table {
	return &schema.Table{
		Name:        "aws_lambda_layer_versions",
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_LayerVersionsListItem.html`,
		Resolver:    fetchLambdaLayerVersions,
		Transform:   transformers.TransformWithStruct(&types.LayerVersionsListItem{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer("lambda"),
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
				Resolver: schema.PathResolver("LayerVersionArn"),
			},
			{
				Name:     "layer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
		},

		Relations: []*schema.Table{
			LayerVersionPolicies(),
		},
	}
}
