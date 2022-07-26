package lightsail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"golang.org/x/sync/errgroup"
)

type DistributionWrapper struct {
	*types.LightsailDistribution
	*lightsail.GetDistributionLatestCacheResetOutput
}

//go:generate cq-gen --resource distributions --config gen.hcl --output .
func Distributions() *schema.Table {
	return &schema.Table{
		Name:          "aws_lightsail_distributions",
		Resolver:      fetchLightsailDistributions,
		Multiplex:     client.AccountMultiplex,
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountFilter,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "able_to_update_bundle",
				Description: "Indicates whether the bundle that is currently applied to your distribution, specified using the distributionName parameter, can be changed to another bundle",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LightsailDistribution.AbleToUpdateBundle"),
			},
			{
				Name:        "alternative_domain_names",
				Description: "The alternate domain names of the distribution",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("LightsailDistribution.AlternativeDomainNames"),
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.Arn"),
			},
			{
				Name:        "bundle_id",
				Description: "The ID of the bundle currently applied to the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.BundleId"),
			},
			{
				Name:        "cache_behavior_settings",
				Description: "An object that describes the cache behavior settings of the distribution",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("LightsailDistribution.CacheBehaviorSettings"),
			},
			{
				Name:        "cache_behaviors",
				Description: "An array of objects that describe the per-path cache behavior of the distribution",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("LightsailDistribution.CacheBehaviors"),
			},
			{
				Name:        "certificate_name",
				Description: "The name of the SSL/TLS certificate attached to the distribution, if any",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.CertificateName"),
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the distribution was created",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("LightsailDistribution.CreatedAt"),
			},
			{
				Name:        "default_cache_behavior",
				Description: "The cache behavior of the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.DefaultCacheBehavior.Behavior"),
			},
			{
				Name:        "domain_name",
				Description: "The domain name of the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.DomainName"),
			},
			{
				Name:        "ip_address_type",
				Description: "The IP address type of the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.IpAddressType"),
			},
			{
				Name:        "is_enabled",
				Description: "Indicates whether the distribution is enabled",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("LightsailDistribution.IsEnabled"),
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.Location.AvailabilityZone"),
			},
			{
				Name:        "name",
				Description: "The name of the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.Name"),
			},
			{
				Name:        "origin_name",
				Description: "The name of the origin resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.Origin.Name"),
			},
			{
				Name:        "origin_protocol_policy",
				Description: "The protocol that your Amazon Lightsail distribution uses when establishing a connection with your origin to pull content",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.Origin.ProtocolPolicy"),
			},
			{
				Name:        "origin_region_name",
				Description: "The AWS Region name of the origin resource",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.Origin.RegionName"),
			},
			{
				Name:        "origin_resource_type",
				Description: "The resource type of the origin resource (eg, Instance)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.Origin.ResourceType"),
			},
			{
				Name:        "origin_public_dns",
				Description: "The public DNS of the origin",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.OriginPublicDNS"),
			},
			{
				Name:        "resource_type",
				Description: "The Lightsail resource type (eg, Distribution)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.ResourceType"),
			},
			{
				Name:        "status",
				Description: "The status of the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.Status"),
			},
			{
				Name:        "support_code",
				Description: "The support code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LightsailDistribution.SupportCode"),
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "cache_reset_create_time",
				Description: "The timestamp of the last cache reset (eg, 147973490917) in Unix time format",
				Type:        schema.TypeTimestamp,
				Resolver:    schema.PathResolver("GetDistributionLatestCacheResetOutput.CreateTime"),
			},
			{
				Name:        "cache_reset_status",
				Description: "The status of the last cache reset",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("GetDistributionLatestCacheResetOutput.Status"),
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
			return diag.WrapError(err)
		}

		errs, ctx := errgroup.WithContext(ctx)
		errs.SetLimit(MAX_GOROUTINES)
		for _, d := range response.Distributions {
			func(d types.LightsailDistribution) {
				errs.Go(func() error {
					return fetchCacheReset(ctx, res, c, d)
				})
			}(d)
		}
		err = errs.Wait()
		if err != nil {
			return diag.WrapError(err)
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
		return diag.WrapError(err)
	}
	res <- DistributionWrapper{&d, resetResp}
	return nil
}
