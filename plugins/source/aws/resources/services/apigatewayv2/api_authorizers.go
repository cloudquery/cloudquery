package apigatewayv2

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApiAuthorizers() *schema.Table {
	tableName := "aws_apigatewayv2_api_authorizers"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Authorizer.html`,
		Resolver:    fetchApigatewayv2ApiAuthorizers,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   client.TransformWithStruct(&types.Authorizer{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "api_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "api_id",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("id"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApiAuthorizerArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
