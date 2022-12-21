package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudfrontCachePolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudfront.ListCachePoliciesInput
	c := meta.(*client.Client)
	s := c.Services()
	svc := s.Cloudfront
	for {
		response, err := svc.ListCachePolicies(ctx, nil)
		if err != nil {
			return err
		}

		if response.CachePolicyList != nil {
			res <- response.CachePolicyList.Items
		}

		if aws.ToString(response.CachePolicyList.NextMarker) == "" {
			break
		}
		config.Marker = response.CachePolicyList.NextMarker
	}
	return nil
}

func resolveCachePolicyARN() schema.ColumnResolver {
	return client.ResolveARNWithAccount(client.CloudfrontService, func(resource *schema.Resource) ([]string, error) {
		return []string{"cache-policy", *resource.Item.(types.CachePolicySummary).CachePolicy.Id}, nil
	})
}
