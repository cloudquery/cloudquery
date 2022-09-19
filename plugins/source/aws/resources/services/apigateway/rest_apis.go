// Code generated by codegen; DO NOT EDIT.

package apigateway

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RestApis() *schema.Table {
	return &schema.Table{
		Name:      "aws_apigateway_rest_apis",
		Resolver:  fetchApigatewayRestApis,
		Multiplex: client.ServiceAccountRegionMultiplexer("apigateway"),
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
			{
				Name:     "api_key_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ApiKeySource"),
			},
			{
				Name:     "binary_media_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("BinaryMediaTypes"),
			},
			{
				Name:     "created_date",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedDate"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "disable_execute_api_endpoint",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DisableExecuteApiEndpoint"),
			},
			{
				Name:     "endpoint_configuration",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EndpointConfiguration"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "minimum_compression_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinimumCompressionSize"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Policy"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
			{
				Name:     "version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Version"),
			},
			{
				Name:     "warnings",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Warnings"),
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
