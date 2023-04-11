package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
)

func Distributions() *schema.Table {
	tableName := "aws_cloudfront_distributions"
	return &schema.Table{
		Name:                tableName,
		Description:         `https://docs.aws.amazon.com/cloudfront/latest/APIReference/API_Distribution.html`,
		Resolver:            fetchCloudfrontDistributions,
		PreResourceResolver: getDistribution,
		Multiplex:           client.ServiceAccountRegionMultiplexer(tableName, "cloudfront"),
		Transform:           transformers.TransformWithStruct(&types.Distribution{}),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudfrontDistributionTags,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
		},
	}
}

func fetchCloudfrontDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudfront.ListDistributionsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudfront
	paginator := cloudfront.NewListDistributionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return err
		}
		res <- page.DistributionList.Items
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
