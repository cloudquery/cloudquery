package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Vaults() *schema.Table {
	tableName := "aws_backup_vaults"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_BackupVaultListMember.html`,
		Resolver:    fetchBackupVaults,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   transformers.TransformWithStruct(&types.BackupVaultListMember{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("BackupVaultArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "access_policy",
				Type:     schema.TypeJSON,
				Resolver: resolveVaultAccessPolicy,
			},
			{
				Name:     "notifications",
				Type:     schema.TypeJSON,
				Resolver: resolveVaultNotifications,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveVaultTags,
			},
		},

		Relations: []*schema.Table{
			VaultRecoveryPoints(),
		},
	}
}
