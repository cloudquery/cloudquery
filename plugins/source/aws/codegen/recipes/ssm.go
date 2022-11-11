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
			Relations: []string{
				`InstanceComplianceItems()`,
				`InstancePatches()`,
			},
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
			ExtraColumns: append(defaultRegionalColumnsPK, []codegen.ColumnDefinition{
				{
					Name:        "name",
					Description: "The parameter name",
					Type:        schema.TypeString,
					Options:     schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
		},
		{
			SubService:  "compliance_summary_items",
			Struct:      &types.ComplianceSummaryItem{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_ComplianceSummaryItem.html",
			SkipFields:  []string{"ComplianceType"},
			ExtraColumns: append(defaultRegionalColumnsPK, []codegen.ColumnDefinition{
				{
					Name:    "compliance_type",
					Type:    schema.TypeString,
					Options: schema.ColumnCreationOptions{PrimaryKey: true},
				},
			}...),
		},
		{
			SubService:  "associations",
			Struct:      &types.Association{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_Association.html",
			SkipFields:  []string{"AssociationId"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "association_id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("AssociationId")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "inventories",
			Struct:      &types.InventoryResultEntity{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryResultEntity.html",
			SkipFields:  []string{"Id"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Id")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "inventory_schemas",
			Struct:      &types.InventoryItemSchema{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_InventoryItemSchema.html",
			SkipFields:  []string{"TypeName", "Version"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "type_name",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("TypeName")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "version",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("Version")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "patch_baselines",
			Struct:      &types.PatchBaselineIdentity{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchBaselineIdentity.html",
			SkipFields:  []string{"BaselineId"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "baseline_id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("BaselineId")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
		{
			SubService:  "instance_patches",
			Struct:      &types.PatchComplianceData{},
			Description: "https://docs.aws.amazon.com/systems-manager/latest/APIReference/API_PatchComplianceData.html",
			SkipFields:  []string{"KBId"},
			ExtraColumns: append(
				defaultRegionalColumnsPK,
				[]codegen.ColumnDefinition{
					{
						Name:     "kb_id",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("KBId")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "ssm"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("ssm")`
	}
	return resources
}
