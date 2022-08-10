package backup

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource global_settings --config gen.hcl --output .
func GlobalSettings() *schema.Table {
	return &schema.Table{
		Name:         "aws_backup_global_settings",
		Resolver:     fetchBackupGlobalSettings,
		Multiplex:    client.ServiceAccountRegionMultiplexer("backup"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id"}},
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "global_settings",
				Description: "The status of the flag isCrossAccountBackupEnabled.",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "last_update_time",
				Description: "The date and time that the flag isCrossAccountBackupEnabled was last updated. This update is in Unix format and Coordinated Universal Time (UTC)",
				Type:        schema.TypeTimestamp,
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBackupGlobalSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Backup
	input := backup.DescribeGlobalSettingsInput{}

	output, err := svc.DescribeGlobalSettings(ctx, &input, func(o *backup.Options) {
		o.Region = c.Region
	})
	if err != nil {
		if client.IgnoreAccessDeniedServiceDisabled(err) || client.IsAWSError(err, "ERROR_9601") /* "Your account is not a member of an organization" */ {
			meta.Logger().Debug("received access denied on DescribeGlobalSettings", "err", err)
			return nil
		}
		if client.IsAWSError(err, "ERROR_2502") /* "Feature Cross Account Backup is not available in current region" */ {
			meta.Logger().Debug("Feature Cross Account Backup is not available in current region on DescribeGlobalSettings", "err", err)
			return nil
		}
		return diag.WrapError(err)
	}
	res <- output
	return nil
}
