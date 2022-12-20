package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

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
