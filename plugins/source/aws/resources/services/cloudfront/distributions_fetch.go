package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchCloudfrontDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudfront.ListDistributionsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudfront
	for {
		response, err := svc.ListDistributions(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.DistributionList.Items

		if aws.ToString(response.DistributionList.Marker) == "" {
			break
		}
		config.Marker = response.DistributionList.Marker
	}
	return nil
}

func getDistribution(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	c := meta.(*client.Client)
	svc := c.Services().Cloudfront

	d := resource.Item.(types.DistributionSummary)

	distribution, err := svc.GetDistribution(ctx, &cloudfront.GetDistributionInput{
		Id: d.Id,
	})
	if err != nil {
		return err
	}
	resource.Item = distribution.Distribution
	return nil
}

func resolveCloudfrontDistributionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	distribution := resource.Item.(*types.Distribution)

	cl := meta.(*client.Client)
	svc := cl.Services().Cloudfront
	response, err := svc.ListTagsForResource(ctx, &cloudfront.ListTagsForResourceInput{
		Resource: distribution.ARN,
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.Tags.Items))
}
