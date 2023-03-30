package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RegionSettings() *schema.Table {
	tableName := "aws_backup_region_settings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeRegionSettings.html`,
		Resolver:    fetchBackupRegionSettings,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   client.TransformWithStruct(&backup.DescribeRegionSettingsOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
