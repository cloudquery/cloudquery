package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RestApis() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_rest_apis",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RestApi.html`,
		Resolver:    fetchApigatewayRestApis,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:  transformers.TransformWithStruct(&types.RestApi{}),
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
				Resolver: resolveApigatewayRestAPIArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			RestApiAuthorizers(),
			RestApiDeployments(),
			RestApiDocumentationParts(),
			RestApiDocumentationVersions(),
			RestApiGatewayResponses(),
			RestApiModels(),
			RestApiRequestValidators(),
			RestApiResources(),
			RestApiStages(),
		},
	}
}
