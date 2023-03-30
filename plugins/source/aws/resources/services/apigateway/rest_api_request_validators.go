package apigateway

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RestApiRequestValidators() *schema.Table {
	tableName := "aws_apigateway_rest_api_request_validators"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_RequestValidator.html`,
		Resolver:    fetchApigatewayRestApiRequestValidators,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "apigateway"),
		Transform:   client.TransformWithStruct(&types.RequestValidator{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(false),
			{
				Name:     "rest_api_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayRestAPIRequestValidatorArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
