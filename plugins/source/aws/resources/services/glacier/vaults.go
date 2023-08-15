package glacier

import (
	"context"

	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"

	"github.com/apache/arrow/go/v13/arrow"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
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
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveGlacierVaultTags,
			},
			{
				Name:       "arn",
				Type:       arrow.BinaryTypes.String,
				Resolver:   schema.PathResolver("VaultARN"),
				PrimaryKey: true,
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
	cl := meta.(*client.Client)
	svc := cl.Services().Glacier
	paginator := glacier.NewListVaultsPaginator(svc, &glacier.ListVaultsInput{})
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *glacier.Options) {
			options.Region = cl.Region
		})
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
	}, func(options *glacier.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		return err
	}
	return resource.Set(c.Name, out.Tags)
}
