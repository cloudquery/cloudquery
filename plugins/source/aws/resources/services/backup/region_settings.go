package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func RegionSettings() *schema.Table {
	return &schema.Table{
		Name:        "aws_backup_region_settings",
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeRegionSettings.html`,
		Resolver:    fetchBackupRegionSettings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("backup"),
		Transform:   transformers.TransformWithStruct(&backup.DescribeRegionSettingsOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}
