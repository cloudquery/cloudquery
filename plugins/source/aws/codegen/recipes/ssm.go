package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SSMResources() []*Resource {
	resources := []*Resource{
		{
			SubService:          "documents",
			Struct:              &types.DocumentDescription{},
			Description:         "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_DocumentDescription.html",
			PreResourceResolver: "getDocument",
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveDocumentARN`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "permissions",
						Type:     schema.TypeJSON,
						Resolver: `resolveDocumentPermission`,
					},
				}...),
		},

		{
			SubService:  "instances",
			Struct:      &types.InstanceInformation{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InstanceInformation.html",
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `resolveInstanceARN`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
			Relations: []string{`InstanceComplianceItems()`},
		},
		{
			SubService:  "instance_compliance_items",
			Struct:      &types.ComplianceItem{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceItem.html",
			SkipFields:  []string{"Id"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Id")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "instance_arn",
						Type:     schema.TypeString,
						Resolver: `resolveInstanceComplianceItemInstanceARN`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},

		{
			SubService:  "parameters",
			Struct:      &types.ParameterMetadata{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ParameterMetadata.html",
			SkipFields:  []string{"Name"},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:        "account_id",
					Description: "The AWS Account ID of the resource",
					Type:        schema.TypeString,
					Resolver:    `client.ResolveAWSAccount`,
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:        "region",
					Description: "The AWS Region of the resource",
					Type:        schema.TypeString,
					Resolver:    `client.ResolveAWSRegion`,
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},

				{
					Name:        "name",
					Description: "The parameter name",
					Type:        schema.TypeString,
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "ssm"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("ssm")`
	}
	return resources
}
