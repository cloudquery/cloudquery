package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/plugin-sdk/codegen"
	"github.com/cloudquery/plugin-sdk/schema"
)

func init() {
	add(combine(&Resource{
		DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn},
		AWSStruct:      &backup.DescribeGlobalSettingsOutput{},
		AWSService:     "Backup",
		Template:       "resource_get",
		ItemsStruct:    &backup.DescribeGlobalSettingsOutput{},
		//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		CustomErrorBlock: `
		if client.IgnoreAccessDeniedServiceDisabled(err) || client.IsAWSError(err, "ERROR_9601") /* "Your account is not a member of an organization" */ {
			meta.Logger().Debug("received access denied on DescribeGlobalSettings", "err", err)
			return nil
		}
		if client.IsAWSError(err, "ERROR_2502") /* "Feature Cross Account Backup is not available in current region" */ {
			meta.Logger().Debug("Feature Cross Account Backup is not available in current region on DescribeGlobalSettings", "err", err)
			return nil
		}
`,
		SkipFields: []string{"ResultMetadata"},
	},
		&Resource{
			DefaultColumns: []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
			AWSStruct:      &backup.DescribeRegionSettingsOutput{},
			AWSService:     "Backup",
			Template:       "resource_get",
			ItemsStruct:    &backup.DescribeRegionSettingsOutput{},
			//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
			SkipFields: []string{"ResultMetadata"},
		},
		parentize(
			&Resource{
				DefaultColumns:       []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
				AWSStruct:            &types.BackupVaultListMember{},
				AWSService:           "Backup",
				Template:             "resource_get",
				CQSubserviceOverride: "vaults",
				ItemsStruct:          &backup.ListBackupVaultsOutput{},
				CustomInputs: []string{
					"\tMaxResults: aws.Int32(1000),", // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupVaults.html
				},
				//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
				ColumnOverrides: map[string]codegen.ColumnDefinition{
					"tags": {
						Type: schema.TypeJSON,
						//					Resolver: ResolverAuto, // TODO enable
					},
				},
			},
			&Resource{
				AWSStruct:            &types.RecoveryPointByBackupVault{},
				Template:             "resource_get",
				CQSubserviceOverride: "recovery_points",
				ItemsStruct:          &backup.ListRecoveryPointsByBackupVaultOutput{},
				CustomInputs: []string{
					"\tMaxResults: aws.Int32(100),",
				},
				ParentFieldName: "BackupVaultName",
				//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
				ColumnOverrides: map[string]codegen.ColumnDefinition{
					"tags": {
						Type: schema.TypeJSON,
						//					Resolver: ResolverAuto, // TODO enable
					},
				},
			},
		),
		parentize(
			&Resource{
				DefaultColumns:       []codegen.ColumnDefinition{AccountIdColumn, RegionColumn},
				AWSStruct:            &backup.GetBackupPlanOutput{},
				AWSService:           "Backup",
				Template:             "resource_list_describe",
				CQSubserviceOverride: "backup_plans",
				PaginatorStruct:      &backup.ListBackupPlansOutput{},
				PaginatorGetStruct:   &backup.GetBackupPlanInput{},
				ItemsStruct:          &backup.GetBackupPlanOutput{},
				CustomInputs: []string{
					"\tMaxResults: aws.Int32(1000),", // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupPlans.html
				},
				//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
				ColumnOverrides: map[string]codegen.ColumnDefinition{
					//"tags": { // TODO enable
					//	Type: schema.TypeJSON,
					//					Resolver: ResolverAuto, // TODO enable
					//},
				},
				Imports:    []string{"github.com/aws/aws-sdk-go-v2/aws"},
				TrimPrefix: "backup_plan_",
			},
			&Resource{
				AWSStruct:            &backup.GetBackupSelectionOutput{},
				Template:             "resource_list_describe",
				CQSubserviceOverride: "backup_selections",
				PaginatorStruct:      &backup.ListBackupSelectionsOutput{},
				PaginatorGetStruct:   &backup.GetBackupSelectionInput{},
				ItemsStruct:          &backup.GetBackupSelectionOutput{},
				CustomInputs: []string{
					"\tMaxResults: aws.Int32(1000),", // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupSelections.html
				},
				ParentFieldName:          "BackupPlanId",
				SkipDescribeParentInputs: true,
				//CreateTableOptions: schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
				ColumnOverrides: map[string]codegen.ColumnDefinition{
					//"tags": { // TODO enable
					//	Type: schema.TypeJSON,
					//					Resolver: ResolverAuto, // TODO enable
					//},
				},
				Imports: []string{"github.com/aws/aws-sdk-go-v2/aws"},
			},
		),
	)...)
}
