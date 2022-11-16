package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApprunnerResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "auto_scaling_configurations",
			Struct:              &types.AutoScalingConfiguration{},
			Description:         "https://docs.aws.amazon.com/apprunner/latest/api/API_AutoScalingConfiguration.html",
			SkipFields:          []string{"AutoScalingConfigurationArn"},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apprunner")`,
			PreResourceResolver: "getAutoScalingConfiguration",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("AutoScalingConfigurationArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveApprunnerTags("AutoScalingConfigurationArn")`,
					},
				}...),
		}, {
			SubService:   "custom_domains",
			Struct:       &types.CustomDomain{},
			Description:  "https://docs.aws.amazon.com/apprunner/latest/api/API_CustomDomain.html",
			Multiplex:    "",
			ExtraColumns: defaultRegionalColumns,
		}, {
			SubService:  "connections",
			Struct:      &types.ConnectionSummary{},
			Description: "https://docs.aws.amazon.com/apprunner/latest/api/API_Connection.html",
			SkipFields:  []string{"ConnectionArn"},
			Multiplex:   `client.ServiceAccountRegionMultiplexer("apprunner")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ConnectionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveApprunnerTags("ConnectionArn")`,
					},
				}...),
		}, {
			SubService:          "observability_configurations",
			Struct:              &types.ObservabilityConfiguration{},
			Description:         "https://docs.aws.amazon.com/apprunner/latest/api/API_ObservabilityConfiguration.html",
			SkipFields:          []string{"ObservabilityConfigurationArn"},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apprunner")`,
			PreResourceResolver: "getObservabilityConfiguration",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ObservabilityConfigurationArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveApprunnerTags("ObservabilityConfigurationArn")`,
					},
				}...),
		}, {
			SubService:   "operations",
			Struct:       &types.OperationSummary{},
			Description:  "https://docs.aws.amazon.com/apprunner/latest/api/API_OperationSummary.html",
			Multiplex:    "",
			ExtraColumns: defaultRegionalColumns,
		}, {
			SubService:          "services",
			Struct:              &types.Service{},
			Description:         "https://docs.aws.amazon.com/apprunner/latest/api/API_Service.html",
			SkipFields:          []string{"ServiceArn"},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apprunner")`,
			PreResourceResolver: "getService",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("ServiceArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveApprunnerTags("ServiceArn")`,
					},
				}...),
			Relations: []string{
				"Operations()",
				"CustomDomains()"},
		},
		{
			SubService:  "vpc_connectors",
			Struct:      &types.VpcConnector{},
			Description: "https://docs.aws.amazon.com/apprunner/latest/api/API_VpcConnector.html",
			SkipFields:  []string{"VpcConnectorArn"},
			Multiplex:   `client.ServiceAccountRegionMultiplexer("apprunner")`,
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("VpcConnectorArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveApprunnerTags("VpcConnectorArn")`,
					},
				}...),
		}, {
			SubService: "vpc_ingress_connections",
			Struct:     &types.VpcIngressConnection{},
			Description: `https://docs.aws.amazon.com/apprunner/latest/api/API_VpcIngressConnection.html

Notes:
- 'account_id' has been renamed to 'source_account_id' to avoid conflict with the 'account_id' column that indicates what account this was synced from.`,
			SkipFields:          []string{"VpcIngressConnectionArn", "AccountId"},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apprunner")`,
			PreResourceResolver: "getVpcIngressConnection",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("VpcIngressConnectionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "source_account_id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("AccountId")`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveApprunnerTags("VpcIngressConnectionArn")`,
					},
				}...),
		},
	}

	for _, r := range resources {
		r.Service = "apprunner"
	}
	return resources
}
