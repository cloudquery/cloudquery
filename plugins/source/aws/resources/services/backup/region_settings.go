package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func RegionSettings() *schema.Table {
	tableName := "aws_backup_region_settings"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeRegionSettings.html`,
		Resolver:    fetchBackupRegionSettings,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   transformers.TransformWithStruct(&backup.DescribeRegionSettingsOutput{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(true),
			client.DefaultRegionColumn(true),
		},
	}
}

func fetchBackupRegionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Backup
	input := backup.DescribeRegionSettingsInput{}

	output, err := svc.DescribeRegionSettings(ctx, &input)
	if err != nil {
		return err
	}
	res <- output
	return nil
}
