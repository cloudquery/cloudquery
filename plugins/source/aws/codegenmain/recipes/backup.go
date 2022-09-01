package recipes

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/plugin-sdk/codegen"
)

var BackupResources = []*Resource{
	{
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
	},
}
