package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func GlobalSettings() *schema.Table {
	return &schema.Table{
		Name:        "aws_backup_global_settings",
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeGlobalSettings.html`,
		Resolver:    fetchBackupGlobalSettings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("backup"),
		Transform:   transformers.TransformWithStruct(&backup.DescribeGlobalSettingsOutput{}),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
