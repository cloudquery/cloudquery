package recipes

import (
	"reflect"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/aws/aws-sdk-go-v2/service/lambda/types"
	lambdaService "github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lambda"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LambdaResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "functions",
			Struct:     &lambda.GetFunctionOutput{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Configuration.FunctionArn")`,
					},
					{
						Name: "policy_revision_id",
						Type: schema.TypeString,
					},
					{
						Name: "policy_document",
						Type: schema.TypeJSON,
					},
					{
						Name: "code_signing_config",
						Type: schema.TypeJSON,
					},
					{
						Name:     "code_repository_type",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Code.RepositoryType")`,
					},
				}...),
			PostResourceResolver: `resolvePolicyCodeSigningConfig`,
			Relations: []string{
				"FunctionEventInvokeConfigs()",
				"FunctionAliases()",
				"FunctionVersions()",
				"FunctionConcurrencyConfigs()",
				"FunctionEventSourceMappings()",
			},
		},
		{
			SubService: "function_event_invoke_configs",
			Struct:     &types.FunctionEventInvokeConfig{},
			SkipFields: []string{"FunctionArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "function_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "function_aliases",
			Struct:     &lambdaService.AliasWrapper{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "function_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "function_versions",
			Struct:     &types.FunctionConfiguration{},
			SkipFields: []string{"FunctionArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "function_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "function_concurrency_configs",
			Struct:     &types.ProvisionedConcurrencyConfigListItem{},
			SkipFields: []string{"FunctionArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "function_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "function_event_source_mappings",
			Struct:     &types.EventSourceMappingConfiguration{},
			SkipFields: []string{"FunctionArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "function_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "layers",
			Struct:     &types.LayersListItem{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("LayerArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{
				"LayerVersions()",
			},
		},
		{
			SubService: "layer_versions",
			Struct:     &types.LayerVersionsListItem{},
			SkipFields: []string{"LayerVersionArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("LayerVersionArn")`,
					},
					{
						Name:     "layer_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
			Relations: []string{
				"LayerVersionPolicies()",
			},
		},
		{
			SubService: "layer_version_policies",
			Struct:     &lambda.GetLayerVersionPolicyOutput{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "layer_version_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "layer_version",
						Type:     schema.TypeInt,
						Resolver: `schema.ParentResourceFieldResolver("version")`,
					},
				}...),
		},
		{
			SubService: "runtimes",
			Struct:     &lambdaService.RuntimeWrapper{},
			SkipFields: []string{"Name"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:    "name",
						Type:    schema.TypeString,
						Options: schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "lambda"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("lambda")`
		structName := reflect.ValueOf(r.Struct).Elem().Type().Name()
		if strings.Contains(structName, "Wrapper") {
			r.UnwrapEmbeddedStructs = true
		}
	}
	return resources
}
