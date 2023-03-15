package glacier

import (
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/cloudquery/plugin-sdk/transformers"
)

func Vaults() *schema.Table {
	tableName := "aws_glacier_vaults"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vaults-get.html`,
		Resolver:    fetchGlacierVaults,
		Transform:   transformers.TransformWithStruct(&types.DescribeVaultOutput{}),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "glacier"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveGlacierVaultTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VaultARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},

		Relations: []*schema.Table{
			VaultAccessPolicies(),
			VaultLockPolicies(),
			VaultNotifications(),
		},
	}
}
