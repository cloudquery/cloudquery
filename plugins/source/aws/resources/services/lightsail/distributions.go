package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/resources/services/lightsail/models"
	"github.com/cloudquery/plugin-sdk/v2/schema"
	"github.com/cloudquery/plugin-sdk/v2/transformers"
	"golang.org/x/sync/errgroup"
)

func Distributions() *schema.Table {
	tableName := "aws_lightsail_distributions"
	return &schema.Table{
		Name:        tableName,
		Description: `https://docs.aws.amazon.com/lightsail/2016-11-28/api-reference/API_GetDistributions.html`,
		Resolver:    fetchLightsailDistributions,
		Transform:   transformers.TransformWithStruct(&models.DistributionWrapper{}, transformers.WithUnwrapAllEmbeddedStructs()),
		Multiplex:   client.ServiceAccountRegionMultiplexer(tableName, "lightsail"),
		Columns: []schema.Column{
			client.DefaultAccountIDColumn(false),
			client.DefaultRegionColumn(false),
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: client.ResolveTags,
			},
		},
	}
}

func fetchLightsailDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	var input lightsail.GetDistributionsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	// No paginator available
	for {
		// Validate the region for this in client/data.json
		response, err := svc.GetDistributions(ctx, &input, func(options *lightsail.Options) {
			// Set region to default global region
			options.Region = "us-east-1"
		})
		if err != nil {
			return err
		}

		errs, ctx := errgroup.WithContext(ctx)
		errs.SetLimit(MaxGoroutines)
		// TODO: Replace with column resolver
		for _, d := range response.Distributions {
			func(d types.LightsailDistribution) {
				errs.Go(func() error {
					return fetchCacheReset(ctx, res, c, d)
				})
			}(d)
		}
		err = errs.Wait()
		if err != nil {
			return err
		}
		if aws.ToString(response.NextPageToken) == "" {
			break
		}
		input.PageToken = response.NextPageToken
	}
	return nil
}

func fetchCacheReset(ctx context.Context, res chan<- any, c *client.Client, d types.LightsailDistribution) error {
	svc := c.Services().Lightsail
	resetInput := lightsail.GetDistributionLatestCacheResetInput{
		DistributionName: d.Name,
	}
	resetResp, err := svc.GetDistributionLatestCacheReset(ctx, &resetInput, func(options *lightsail.Options) {
		// Set region to default global region
		options.Region = "us-east-1"
	})
	if err != nil && !c.IsNotFoundError(err) {
		return err
	}
	res <- models.DistributionWrapper{LightsailDistribution: &d, LatestCacheReset: resetResp}
	return nil
}
