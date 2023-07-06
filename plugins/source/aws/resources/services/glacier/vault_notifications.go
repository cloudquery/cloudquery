package glacier

import (
	"context"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
)

func vaultNotifications() *schema.Table {
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
				Name:       "vault_arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.ParentColumnResolver("arn"),
				PrimaryKey: true,
			},
		},
	}
}

func fetchGlacierVaultNotifications(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glacier
	p := parent.Item.(types.DescribeVaultOutput)

	response, err := svc.GetVaultNotifications(ctx, &glacier.GetVaultNotificationsInput{
		VaultName: p.VaultName,
	}, func(options *glacier.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	res <- response.VaultNotificationConfig
	return nil
}
