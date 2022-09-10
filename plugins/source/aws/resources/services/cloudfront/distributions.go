package cloudfront

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func CloudfrontDistributions() *schema.Table {
	return &schema.Table{
		Name:        "aws_cloudfront_distributions",
		Description: "A summary of the information about a CloudFront distribution.",
		Resolver:    fetchCloudfrontDistributions,
		Multiplex:   client.AccountMultiplex,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveCloudfrontDistributionTags,
			},
			{
				Name:            "arn",
				Description:     "The ARN (Amazon Resource Name) for the distribution",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "distribution_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DistributionConfig"),
			},
			{
				Name:        "domain_name",
				Description: "The domain name corresponding to the distribution, for example, d111111abcdef8.cloudfront.net.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The identifier for the distribution",
				Type:        schema.TypeString,
			},
			{
				Name:        "in_progress_invalidation_batches",
				Description: "The number of invalidation batches currently in progress.",
				Type:        schema.TypeInt,
			},
			{
				Name:        "last_modified_time",
				Description: "The date and time the distribution was last modified.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "status",
				Description: "This response element indicates the current status of the distribution",
				Type:        schema.TypeString,
			},
			{
				Name:     "active_trusted_key_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveTrustedKeyGroups"),
			},
			{
				Name:     "active_trusted_signers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveTrustedSigners"),
			},
			{
				Name:        "alias_icp_recordals",
				Description: "AWS services in China customers must file for an Internet Content Provider (ICP) recordal if they want to serve content publicly on an alternate domain name, also known as a CNAME, that they've added to CloudFront",
				Type:        schema.TypeJSON,
				Resolver:    resolveCloudfrontDistributionsAliasIcpRecordals,
			},
			{
				Name:        "config",
				Description: "Distribution configuration",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("DistributionConfig"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchCloudfrontDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config cloudfront.ListDistributionsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudfront
	for {
		response, err := svc.ListDistributions(ctx, &config)
		if err != nil {
			return err
		}
		for _, d := range response.DistributionList.Items {
			distribution, err := svc.GetDistribution(ctx, &cloudfront.GetDistributionInput{
				Id: d.Id,
			}, func(options *cloudfront.Options) {
				options.Region = c.Region
			})
			if err != nil {
				return err
			}
			res <- *distribution.Distribution
		}

		if aws.ToString(response.DistributionList.Marker) == "" {
			break
		}
		config.Marker = response.DistributionList.Marker
	}
	return nil
}
func resolveCloudfrontDistributionTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	distribution := resource.Item.(types.Distribution)

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
func resolveCloudfrontDistributionsActiveTrustedKeyGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	distribution := resource.Item.(types.Distribution)
	if distribution.ActiveTrustedKeyGroups == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, k := range distribution.ActiveTrustedKeyGroups.Items {
		j[*k.KeyGroupId] = k.KeyPairIds.Items
	}
	return resource.Set(c.Name, j)
}
func resolveCloudfrontDistributionsActiveTrustedSigners(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	distribution := resource.Item.(types.Distribution)
	if distribution.ActiveTrustedSigners == nil {
		return nil
	}
	j := map[string]interface{}{}
	for _, k := range distribution.ActiveTrustedSigners.Items {
		j[*k.AwsAccountNumber] = k.KeyPairIds.Items
	}
	return resource.Set(c.Name, j)
}
func resolveCloudfrontDistributionsAliasIcpRecordals(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	distribution := resource.Item.(types.Distribution)
	j := map[string]interface{}{}
	for _, a := range distribution.AliasICPRecordals {
		j[*a.CNAME] = a.ICPRecordalStatus
	}
	return resource.Set(c.Name, j)
}
