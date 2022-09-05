package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func ResolveCachePolicyArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	if err := client.ResolveARNWithAccount(client.CloudfrontService, func(resource *schema.Resource) ([]string, error) {
		return []string{"cache-policy", *resource.Item.(types.CachePolicySummary).CachePolicy.Id}, nil
	})(ctx, meta, resource, c); err != nil {
		return diag.WrapError(err)
	}
	return nil
}
