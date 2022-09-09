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
				Name:        "caller_reference",
				Description: "A unique value (for example, a date-time stamp) that ensures that the request can't be replayed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.CallerReference"),
			},
			{
				Name:     "distribution_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DistributionConfig"),
			},
			{
				Name:        "cache_behavior_field_level_encryption_id",
				Description: "The value of ID for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for the default cache behavior.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.FieldLevelEncryptionId"),
			},
			{
				Name:          "cache_behavior_realtime_log_config_arn",
				Description:   "The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DistributionConfig.DefaultCacheBehavior.RealtimeLogConfigArn"),
				IgnoreInTests: true,
			},
			{
				Name:        "cache_behavior_smooth_streaming",
				Description: "Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify true; if not, specify false",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.SmoothStreaming"),
			},
			{
				Name:        "default_root_object",
				Description: "The object that you want CloudFront to request from your origin (for example, index.html) when a viewer requests the root URL for your distribution (http://www.example.com) instead of an object in your distribution (http://www.example.com/product-description.html)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultRootObject"),
			},
			{
				Name:        "http_version",
				Description: "(Optional) Specify the maximum HTTP version that you want viewers to use to communicate with CloudFront",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.HttpVersion"),
			},
			{
				Name:        "logging_prefix",
				Description: "An optional string that you want CloudFront to prefix to the access log filenames for this distribution, for example, myprefix/",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.Prefix"),
			},
			{
				Name:        "geo_restriction_type",
				Description: "The method that you want to use to restrict distribution of your content by country:  * none: No geo restriction is enabled, meaning access to content is not restricted by client geo location.  * blacklist: The Location elements specify the countries in which you don't want CloudFront to distribute your content.  * whitelist: The Location elements specify the countries in which you want CloudFront to distribute your content.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Restrictions.GeoRestriction.RestrictionType"),
			},
			{
				Name:        "geo_restrictions",
				Description: "A complex type that contains a Location element for each country in which you want CloudFront either to distribute your content (whitelist) or not distribute your content (blacklist)",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.Restrictions.GeoRestriction.Items"),
			},
			{
				Name:          "viewer_certificate_acm_certificate_arn",
				Description:   "If the distribution uses Aliases (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in AWS Certificate Manager (ACM) (https://docs.aws.amazon.com/acm/latest/userguide/acm-overview.html), provide the Amazon Resource Name (ARN) of the ACM certificate",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DistributionConfig.ViewerCertificate.ACMCertificateArn"),
				IgnoreInTests: true,
			},
			{
				Name:          "viewer_certificate_iam_certificate_id",
				Description:   "If the distribution uses Aliases (alternate domain names or CNAMEs) and the SSL/TLS certificate is stored in AWS Identity and Access Management (AWS IAM) (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_server-certs.html), provide the ID of the IAM certificate",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DistributionConfig.ViewerCertificate.IAMCertificateId"),
				IgnoreInTests: true,
			},
			{
				Name:        "viewer_certificate_minimum_protocol_version",
				Description: "If the distribution uses Aliases (alternate domain names or CNAMEs), specify the security policy that you want CloudFront to use for HTTPS connections with viewers",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.MinimumProtocolVersion"),
			},
			{
				Name:        "viewer_certificate_ssl_support_method",
				Description: "If the distribution uses Aliases (alternate domain names or CNAMEs), specify which viewers the distribution accepts HTTPS connections from.  * sni-only – The distribution accepts HTTPS connections from only viewers that support server name indication (SNI) (https://en.wikipedia.org/wiki/Server_Name_Indication). This is recommended",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.SSLSupportMethod"),
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
