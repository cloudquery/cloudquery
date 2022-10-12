package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/glacier"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchGlacierVaults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Glacier
	input := glacier.ListVaultsInput{}
	for {
		output, err := svc.ListVaults(ctx, &input)
		if err != nil {
			return err
		}
		res <- output.VaultList

		if aws.ToString(output.Marker) == "" {
			break
		}
		input.Marker = output.Marker
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
