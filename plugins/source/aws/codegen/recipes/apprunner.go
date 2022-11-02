package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/apprunner/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func ApprunnerResources() []*Resource {
	resources := []*Resource{
		{
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
				}...),
			Relations: []string{
				"Operations()",
				"CustomDomains()"},
		}, {
			SubService:          "auto_scaling_configuration",
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
				}...),
		},
		{
			SubService:  "vpc_connector",
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
				}...),
		}, {
			SubService:          "vpc_ingress_connection",
			Struct:              &types.VpcIngressConnection{},
			Description:         "https://docs.aws.amazon.com/apprunner/latest/api/API_VpcIngressConnection.html",
			SkipFields:          []string{"VpcIngressConnectionArn"},
			Multiplex:           `client.ServiceAccountRegionMultiplexer("apprunner")`,
			PreResourceResolver: "getVpcIngressConnection",
			ExtraColumns: append(
				[]codegen.ColumnDefinition{
					{
						Name:     "region",
						Type:     schema.TypeString,
						Resolver: "client.ResolveAWSRegion",
					},
				},
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("VpcIngressConnectionArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		}, {
			SubService:   "operations",
			Struct:       &types.OperationSummary{},
			Description:  "https://docs.aws.amazon.com/apprunner/latest/api/API_OperationSummary.html",
			Multiplex:    "",
			ExtraColumns: defaultRegionalColumns,
		}, {
			SubService:   "custom_domains",
			Struct:       &types.CustomDomain{},
			Description:  "https://docs.aws.amazon.com/apprunner/latest/api/API_CustomDomain.html",
			Multiplex:    "",
			ExtraColumns: defaultRegionalColumns,
		},
	}

	for _, r := range resources {
		r.Service = "apprunner"
	}
	return resources
}
