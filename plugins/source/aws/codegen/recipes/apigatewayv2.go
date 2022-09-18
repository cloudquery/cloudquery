package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigatewayv2/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func APIGatewayV2Resources() []*Resource {
	resources := []*Resource{
		{
			SubService: "apis",
			Struct:     &types.Api{},
			SkipFields: []string{"ApiId"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiArn()`,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ApiId")`,
					},
				}...),
			Relations: []string{
				"ApiAuthorizers()",
				"ApiDeployments()",
				"ApiIntegrations()",
				"ApiModels()",
				"ApiRoutes()",
				"ApiStages()",
			},
		},
		{
			SubService: "api_authorizers",
			Struct:     &types.Authorizer{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiAuthorizerArn()`,
					},
				}...),
		},
		{
			SubService: "api_deployments",
			Struct:     &types.Deployment{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiDeploymentArn()`,
					},
				}...),
		},
		{
			SubService: "api_integrations",
			Struct:     &types.Integration{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiIntegrationArn()`,
					},
				}...),
			Relations: []string{
				"ApiIntegrationResponses()",
			},
		},
		{
			SubService: "api_integration_responses",
			Struct:     &types.IntegrationResponse{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_integration_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "integration_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("integration_id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiIntegrationResponseArn()`,
					},
				}...),
		},
		{
			SubService: "api_models",
			Struct:     &types.Model{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiModelArn()`,
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayv2apiModelModelTemplate`,
					},
				}...),
		},
		{
			SubService: "api_routes",
			Struct:     &types.Route{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiRouteArn()`,
					},
				}...),
			Relations: []string{
				"ApiRouteResponses()",
			},
		},
		{
			SubService: "api_route_responses",
			Struct:     &types.RouteResponse{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_route_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "route_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("route_id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiRouteResponseArn()`,
					},
				}...),
		},
		{
			SubService: "api_stages",
			Struct:     &types.Stage{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "api_id",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("id")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApiStageArn()`,
					},
				}...),
		},
		{
			SubService: "domain_names",
			Struct:     &types.DomainName{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveDomainNameArn()`,
					},
				}...),
			Relations: []string{
				"DomainNameRestApiMappings()",
			},
		},
		{
			SubService: "domain_name_rest_api_mappings",
			Struct:     &types.ApiMapping{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "domain_name_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveDomainNameRestApiMappingArn()`,
					},
				}...),
		},
		{
			SubService: "vpc_links",
			Struct:     &types.VpcLink{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveVpcLinkArn()`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "apigatewayv2"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("apigateway")`
	}
	return resources
}
