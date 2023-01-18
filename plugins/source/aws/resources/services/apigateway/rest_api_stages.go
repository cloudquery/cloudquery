package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RestApiStages() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_rest_api_stages",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html`,
		Resolver:    fetchApigatewayRestApiStages,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:   transformers.TransformWithStruct(&types.Stage{}),
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
				Name:     "rest_api_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayRestAPIStageArn,
			},
		},
	}
}
