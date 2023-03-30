package glacier

import (
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func VaultNotifications() *schema.Table {
	tableName := "aws_glacier_vault_notifications"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-get.html`,
		Resolver:    fetchGlacierVaultNotifications,
		Transform:   transformers.TransformWithStruct(&types.VaultNotificationConfig{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glacier"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "vault_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}
