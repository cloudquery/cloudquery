package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApiGatewayV2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "apis",
			Struct:              &types.Api{},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apigatewayv2")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:          "vpc_links",
			Struct:              &types.VpcLink{},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apigatewayv2")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcLinkArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:          "domain_name_api_mappings",
			Struct:              &types.ApiMapping{},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apigatewayv2")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveDomainNameApiMappingArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:          "domain_names",
			Struct:              &types.DomainName{},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apigatewayv2")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveDomainNameArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "apigatewayv2"
	}
	return resources
}
