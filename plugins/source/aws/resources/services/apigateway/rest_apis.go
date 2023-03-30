package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RestApis() *schema.Table {
	tableName := "aws_apigateway_rest_apis"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RestApi.html`,
		Resolver:    fetchApigatewayRestApis,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   client.TransformWithStruct(&types.RestApi{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
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
