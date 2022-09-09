package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"golang.org/x/sync/errgroup"
)

func Distributions() *schema.Table {
	return &schema.Table{
		Name:          "aws_lightsail_distributions",
		Resolver:      fetchLightsailDistributions,
		Multiplex:     client.AccountMultiplex,
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:     "lightsail_distribution",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("LightsailDistribution"),
			},
			{
				Name:            "arn",
				Description:     "The Amazon Resource Name (ARN) of the distribution",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("LightsailDistribution.Arn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:     "get_distribution_latest_cache_reset_output",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("GetDistributionLatestCacheResetOutput"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var input lightsail.GetDistributionsInput
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	for {
		response, err := svc.GetDistributions(ctx, &input, func(options *lightsail.Options) {
			// Set region to default global region
			options.Region = "us-east-1"
		})
		if err != nil {
			return err
		}

		errs, ctx := errgroup.WithContext(ctx)
		errs.SetLimit(MaxGoroutines)
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

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func fetchCacheReset(ctx context.Context, res chan<- interface{}, c *client.Client, d types.LightsailDistribution) error {
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
	res <- DistributionWrapper{&d, resetResp}
	return nil
}
