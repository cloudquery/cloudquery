package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func APIGatewayResources() []*Resource {
	resources := []*Resource{
		{
			SubService:  "api_keys",
			Struct:      &types.ApiKey{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_ApiKey.html",
			SkipFields:  []string{},
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
			SubService:  "client_certificates",
			Struct:      &types.ClientCertificate{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_ClientCertificate.html",
			SkipFields:  []string{},
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
			SubService:  "domain_names",
			Struct:      &types.DomainName{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_DomainName.html",
			SkipFields:  []string{},
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
			SubService:  "domain_name_base_path_mappings",
			Struct:      &types.BasePathMapping{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_BasePathMapping.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "domain_name_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayDomainNameBasePathMappingArn`,
					},
				}...),
		},
		{
			SubService:  "rest_apis",
			Struct:      &types.RestApi{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_RestApi.html",
			SkipFields:  []string{},
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
			SubService:  "rest_api_authorizers",
			Struct:      &types.Authorizer{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Authorizer.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIAuthorizerArn`,
					},
				}...),
		},
		{
			SubService:  "rest_api_deployments",
			Struct:      &types.Deployment{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Deployment.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIDeploymentArn`,
					},
				}...),
		},
		{
			SubService:  "rest_api_documentation_parts",
			Struct:      &types.DocumentationPart{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationPart.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIDocumentationPartArn`,
					},
				}...),
		},
		{
			SubService:  "rest_api_documentation_versions",
			Struct:      &types.DocumentationVersion{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_DocumentationVersion.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIDocumentationVersionArn`,
					},
				}...),
		},
		{
			SubService:  "rest_api_gateway_responses",
			Struct:      &types.GatewayResponse{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_GatewayResponse.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIGatewayResponseArn`,
					},
				}...),
		},
		{
			SubService:  "rest_api_models",
			Struct:      &types.Model{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Model.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
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
			SubService:  "rest_api_request_validators",
			Struct:      &types.RequestValidator{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_RequestValidator.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIRequestValidatorArn`,
					},
				}...),
		},
		{
			SubService:  "rest_api_resources",
			Struct:      &types.Resource{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Resource.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIResourceArn`,
					},
				}...),
		},
		{
			SubService:  "rest_api_stages",
			Struct:      &types.Stage{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_Stage.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "rest_api_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayRestAPIStageArn`,
					},
				}...),
		},
		{
			SubService:  "usage_plans",
			Struct:      &types.UsagePlan{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlan.html",
			SkipFields:  []string{},
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
			SubService:  "usage_plan_keys",
			Struct:      &types.UsagePlanKey{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_UsagePlanKey.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "usage_plan_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveApigatewayUsagePlanKeyArn`,
					},
				}...),
		},
		{
			SubService:  "vpc_links",
			Struct:      &types.VpcLink{},
			Description: "https://docs.aws.amazon.com/apigateway/latest/api/API_VpcLink.html",
			SkipFields:  []string{},
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
