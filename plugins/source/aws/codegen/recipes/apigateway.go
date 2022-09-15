package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func APIGatewayResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "api_keys",
			Struct:     &types.ApiKey{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayAPIKeyArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService: "client_certificates",
			Struct:     &types.ClientCertificate{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayClientCertificateArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
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
						Resolver: `resolveApigatewayDomainNameArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"DomainNameBasePathMappings()",
			},
		},
		{
			SubService: "domain_name_base_path_mappings",
			Struct:     &types.BasePathMapping{},
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
						Resolver: `resolveApigatewayDomainNameBasePathMappingArn`,
					},
				}...),
		},
		{
			SubService: "rest_apis",
			Struct:     &types.RestApi{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIArn`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"RestApiAuthorizers()",
				"RestApiDeployments()",
				"RestApiDocumentationParts()",
				"RestApiDocumentationVersions()",
				"RestApiGatewayResponses()",
				"RestApiModels()",
				"RestApiRequestValidators()",
				"RestApiResources()",
				"RestApiStages()",
			},
		},
		{
			SubService: "rest_api_authorizers",
			Struct:     &types.Authorizer{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIAuthorizerArn`,
					},
				}...),
		},
		{
			SubService: "rest_api_deployments",
			Struct:     &types.Deployment{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIDeploymentArn`,
					},
				}...),
		},
		{
			SubService: "rest_api_documentation_parts",
			Struct:     &types.DocumentationPart{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIDocumentationPartArn`,
					},
				}...),
		},
		{
			SubService: "rest_api_documentation_versions",
			Struct:     &types.DocumentationVersion{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIDocumentationVersionArn`,
					},
				}...),
		},
		{
			SubService: "rest_api_gateway_responses",
			Struct:     &types.GatewayResponse{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIGatewayResponseArn`,
					},
				}...),
		},
		{
			SubService: "rest_api_models",
			Struct:     &types.Model{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIModelArn`,
					},
					{
						Name:     "model_template",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIModelModelTemplate`,
					},
				}...),
		},
		{
			SubService: "rest_api_request_validators",
			Struct:     &types.RequestValidator{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIRequestValidatorArn`,
					},
				}...),
		},
		{
			SubService: "rest_api_resources",
			Struct:     &types.Resource{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIResourceArn`,
					},
				}...),
		},
		{
			SubService: "rest_api_stages",
			Struct:     &types.Stage{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIStageArn`,
					},
				}...),
		},
		{
			SubService: "usage_plans",
			Struct:     &types.UsagePlan{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayUsagePlanArn`,
					},
				}...),
			Relations: []string{
				"UsagePlanKeys()",
			},
		},
		{
			SubService: "usage_plan_keys",
			Struct:     &types.UsagePlanKey{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "usage_plan_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayUsagePlanKeyArn`,
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
						Resolver: `resolveApigatewayVpcLinkArn`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "apigateway"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("apigateway")`
	}
	return resources
}
