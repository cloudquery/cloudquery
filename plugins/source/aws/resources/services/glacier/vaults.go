package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
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
			vaultAccessPolicies(),
			vaultLockPolicies(),
			vaultNotifications(),
		},
	}
}

func fetchGlacierVaults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	c := meta.(*client.Client)
	svc := c.Services().Glacier
	paginator := glacier.NewListVaultsPaginator(svc, &glacier.ListVaultsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.VaultList
	}
	return nil
}

func resolveGlacierVaultTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glacier
	it := resource.Item.(types.DescribeVaultOutput)
	out, err := svc.ListTagsForVault(ctx, &glacier.ListTagsForVaultInput{
		VaultName: it.VaultName,
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
