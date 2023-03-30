package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func GlobalSettings() *schema.Table {
	tableName := "aws_backup_global_settings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeGlobalSettings.html`,
		Resolver:    fetchBackupGlobalSettings,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   client.TransformWithStruct(&backup.DescribeGlobalSettingsOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
