// Code generated by codegen; DO NOT EDIT.

package apigateway

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RestApiModels() *schema.Table {
	return &schema.Table{
		Name:        "aws_apigateway_rest_api_models",
		Description: `https://docs.aws.amazon.com/apigateway/latest/api/API_Model.html`,
		Resolver:    fetchApigatewayRestApiModels,
		Multiplex:   client.ServiceAccountRegionMultiplexer("apigateway"),
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
				Resolver: resolveApigatewayRestAPIModelArn,
			},
			{
				Name:     "model_template",
				Type:     schema.TypeString,
				Resolver: resolveApigatewayRestAPIModelModelTemplate,
			},
			{
				Name:     "content_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ContentType"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "schema",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Schema"),
			},
		},
	}
}
