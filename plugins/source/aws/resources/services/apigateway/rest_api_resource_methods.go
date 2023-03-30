package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func restApiResourceMethods() *schema.Table {
	tableName := "aws_apigateway_rest_api_resource_methods"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Method.html`,
		Resolver:    fetchApigatewayRestApiResourceMethods,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   client.TransformWithStruct(&apigateway.GetMethodOutput{}, transformers.WithSkipFields("ResultMetadata")),
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
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayRestAPIResourceMethodArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
		Relations: []*schema.Table{
			restApiResourceMethodIntegrations(),
		},
	}
}
