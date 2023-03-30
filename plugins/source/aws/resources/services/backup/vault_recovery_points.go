package backup

import (
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func VaultRecoveryPoints() *schema.Table {
	tableName := "aws_backup_vault_recovery_points"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/aws-backup/latest/devguide/API_RecoveryPointByBackupVault.html`,
		Resolver:    fetchBackupVaultRecoveryPoints,
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "backup"),
		Transform:   client.TransformWithStruct(&types.RecoveryPointByBackupVault{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "vault_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RecoveryPointArn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveRecoveryPointTags,
			},
		},
	}
}
