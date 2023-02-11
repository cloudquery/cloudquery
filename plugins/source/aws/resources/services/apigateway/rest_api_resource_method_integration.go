package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func restApiResourceMethodIntegrations() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_rest_api_resource_method_integrations",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Integration.html`,
		Resolver:    fetchApigatewayRestApiResourceMethodIntegration,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
		Transform:   transformers.TransformWithStruct(&apigateway.GetIntegrationOutput{}, transformers.WithSkipFields("ResultMetadata")),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "rest_api_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("rest_api_arn"),
			},
			{
				Name:     "resource_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("resource_arn"),
			},
			{
				Name:     "method_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayRestAPIResourceMethodIntegrationArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{},
	}
}
