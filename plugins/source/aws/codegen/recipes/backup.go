package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BackupResources() []*Resource {
	resources := []*Resource{
		{
			SubService: "global_settings",
			Struct:     &backup.DescribeGlobalSettingsOutput{},
			SkipFields: []string{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService: "plans",
			Struct:     &backup.GetBackupPlanOutput{},
			SkipFields: []string{"BackupPlanArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("BackupPlanArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolvePlanTags`,
					},
				}...),
			Relations: []string{
				"PlanSelections()",
			},
		},
		{
			SubService: "plan_selections",
			Struct:     &backup.GetBackupSelectionOutput{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "plan_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
				}...),
		},
		{
			SubService: "region_settings",
			Struct:     &backup.DescribeRegionSettingsOutput{},
			SkipFields: []string{},
			ExtraColumns: []codegen.ColumnDefinition{
				{
					Name:     "account_id",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSAccount`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
				{
					Name:     "region",
					Type:     schema.TypeString,
					Resolver: `client.ResolveAWSRegion`,
					Options:  schema.ColumnCreationOptions{PrimaryKey: true},
				},
			},
		},
		{
			SubService: "vaults",
			Struct:     &types.BackupVaultListMember{},
			SkipFields: []string{"BackupVaultArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("BackupVaultArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
					},
					{
						Name:          "access_policy",
						Type:          schema.TypeJSON,
						Resolver:      `resolveVaultAccessPolicy`,
						IgnoreInTests: true,
					},
					{
						Name:          "notifications",
						Type:          schema.TypeJSON,
						Resolver:      `resolveVaultNotifications`,
						IgnoreInTests: true,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveVaultTags`,
					},
				}...),
			Relations: []string{
				"VaultRecoveryPoints()",
			},
		},
		{
			SubService: "vault_recovery_points",
			Struct:     &types.RecoveryPointByBackupVault{},
			SkipFields: []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "vault_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentResourceFieldResolver("arn")`,
					},
					{
						Name:     "tags",
						Type:     schema.TypeJSON,
						Resolver: `resolveRecoveryPointTags`,
					},
				}...),
		},
	}

	// set default values
	for _, r := range resources {
		r.Service = "backup"
		r.Multiplex = `client.ServiceAccountRegionMultiplexer("backup")`
	}
	return resources
}
