package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchBackupGlobalSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Backup
	input := backup.DescribeGlobalSettingsInput{}

	output, err := svc.DescribeGlobalSettings(ctx, &input)
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) || client.IsAWSError(err, "ERROR_9601") /* "Your account is not a member of an organization" */ {
			meta.Logger().Debug().Err(err).Msg("received access denied on DescribeGlobalSettings")
			return nil
		}
		if client.IsAWSError(err, "ERROR_2502") /* "Feature Cross Account Backup is not available in current region" */ {
			meta.Logger().Debug().Err(err).Msg("Feature Cross Account Backup is not available in current region on DescribeGlobalSettings")
			return nil
		}
		return err
	}
	res <- output
	return nil
}
