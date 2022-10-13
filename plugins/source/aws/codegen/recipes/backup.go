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
			SubService:  "global_settings",
			Description: "https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeGlobalSettings.html",
			Struct:      &backup.DescribeGlobalSettingsOutput{},
			SkipFields:  []string{},
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
			SubService:          "plans",
			Description:         "https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupPlan.html",
			Struct:              &backup.GetBackupPlanOutput{},
			SkipFields:          []string{"BackupPlanArn"},
			PreResourceResolver: "getPlan",
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
			SubService:  "plan_selections",
			Description: "https://docs.aws.amazon.com/aws-backup/latest/devguide/API_GetBackupSelection.html",
			Struct:      &backup.GetBackupSelectionOutput{},
			SkipFields:  []string{},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "plan_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
				}...),
		},
		{
			SubService:  "region_settings",
			Description: "https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeRegionSettings.html",
			Struct:      &backup.DescribeRegionSettingsOutput{},
			SkipFields:  []string{},
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
			SubService:  "vaults",
			Struct:      &types.BackupVaultListMember{},
			Description: "https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BackupVaultListMember.html",
			SkipFields:  []string{"BackupVaultArn"},
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
			SubService:  "vault_recovery_points",
			Struct:      &types.RecoveryPointByBackupVault{},
			Description: "https://docs.aws.amazon.com/aws-backup/latest/devguide/API_RecoveryPointByBackupVault.html",
			SkipFields:  []string{"RecoveryPointArn"},
			ExtraColumns: append(
				defaultRegionalColumns,
				[]codegen.ColumnDefinition{
					{
						Name:     "vault_arn",
						Type:     schema.TypeString,
						Resolver: `schema.ParentColumnResolver("arn")`,
					},
					{
						Name:     "arn",
						Type:     schema.TypeString,
						Resolver: `schema.PathResolver("RecoveryPointArn")`,
						Options:  schema.ColumnCreationOptions{PrimaryKey: true},
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
