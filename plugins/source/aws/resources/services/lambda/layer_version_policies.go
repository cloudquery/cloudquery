package lambda

import (
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func LayerVersionPolicies() *schema.Table {
	return &schema.Table{
		Name:        "aws_lambda_layer_version_policies",
		Description: `https://docs.aws.amazon.com/lambda/latest/dg/API_GetLayerVersionPolicy.html`,
		Resolver:    fetchLambdaLayerVersionPolicies,
		Transform:   transformers.TransformWithStruct(&lambda.GetLayerVersionPolicyOutput{}),
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
				Name:     "layer_version_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "layer_version",
				Type:     schema.TypeInt,
				Resolver: schema.ParentColumnResolver("version"),
			},
		},
	}
}
