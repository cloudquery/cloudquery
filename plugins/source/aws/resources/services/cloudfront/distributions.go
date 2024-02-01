package cloudfront

import (
	"context"

	"github.com/apache/arrow/go/v15/arrow"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	sdkTypes "github.com/cloudquery/plugin-sdk/v4/types"
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
				Type:     sdkTypes.ExtensionTypes.JSON,
				Resolver: resolveCloudfrontDistributionTags,
			},
			{
				Name:                "arn",
				Type:                arrow.BinaryTypes.String,
				Resolver:            schema.PathResolver("ARN"),
				PrimaryKeyComponent: true,
			},
		},
	}
}

func fetchCloudfrontDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var config cloudfront.ListDistributionsInput
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudfront).Cloudfront
	paginator := cloudfront.NewListDistributionsPaginator(svc, &config)
	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx, func(options *cloudfront.Options) {
			options.Region = cl.Region
		})
		if err != nil {
			return err
		}
		res <- page.DistributionList.Items
	}
	return nil
}

func getDistribution(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	cl := meta.(*client.Client)
	svc := cl.Services(client.AWSServiceCloudfront).Cloudfront

	d := resource.Item.(types.DistributionSummary)

	distribution, err := svc.GetDistribution(ctx, &cloudfront.GetDistributionInput{
		Id: d.Id,
	}, func(options *cloudfront.Options) {
		options.Region = cl.Region
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
	svc := cl.Services(client.AWSServiceCloudfront).Cloudfront
	response, err := svc.ListTagsForResource(ctx, &cloudfront.ListTagsForResourceInput{
		Resource: distribution.ARN,
	}, func(options *cloudfront.Options) {
		options.Region = cl.Region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, client.TagsToMap(response.Tags.Items))
}
