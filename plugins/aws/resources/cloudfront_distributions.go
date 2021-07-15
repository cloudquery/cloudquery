package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func CloudfrontDistributions() *schema.Table {
	return &schema.Table{
		Name:         "aws_cloudfront_distributions",
		Resolver:     fetchCloudfrontDistributions,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
			},
			{
				Name:     "aliases_items",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Aliases.Items"),
			},
			{
				Name: "comment",
				Type: schema.TypeString,
			},
			{
				Name:     "cache_behaviour_target_origin_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TargetOriginId"),
			},
			{
				Name:     "cache_behaviour_viewer_protocol_policy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ViewerProtocolPolicy"),
			},
			{
				Name:     "cache_behaviour_allowed_methods",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.AllowedMethods.Items"),
			},
			{
				Name:     "cache_behaviour_allowed_methods_cached_methods",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.AllowedMethods.CachedMethods.Items"),
			},
			{
				Name:     "cache_behaviour_cache_policy_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.CachePolicyId"),
			},
			{
				Name:     "cache_behaviour_compress",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.Compress"),
			},
			{
				Name:     "cache_behaviour_default_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DefaultCacheBehavior.DefaultTTL"),
			},
			{
				Name:     "cache_behaviour_field_level_encryption_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.FieldLevelEncryptionId"),
			},
			{
				Name:     "cache_behaviour_forwarded_values_cookies_forward",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.Cookies.Forward"),
			},
			{
				Name:     "cache_behaviour_forwarded_values_cookies_white_listed_names",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.Cookies.WhitelistedNames.Items"),
			},
			{
				Name:     "cache_behaviour_forwarded_values_query_string",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.QueryString"),
			},
			{
				Name:     "cache_behaviour_forwarded_values_headers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.Headers.Items"),
			},
			{
				Name:     "cache_behaviour_forwarded_values_query_string_cache_keys",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.ForwardedValues.QueryStringCacheKeys.Items"),
			},
			{
				Name:     "cache_behaviour_max_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DefaultCacheBehavior.MaxTTL"),
			},
			{
				Name:     "cache_behaviour_min_ttl",
				Type:     schema.TypeBigInt,
				Resolver: schema.PathResolver("DefaultCacheBehavior.MinTTL"),
			},
			{
				Name:     "cache_behaviour_origin_request_policy_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.OriginRequestPolicyId"),
			},
			{
				Name:     "cache_behaviour_realtime_log_config_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultCacheBehavior.RealtimeLogConfigArn"),
			},
			{
				Name:     "cache_behaviour_smooth_streaming",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.SmoothStreaming"),
			},
			{
				Name:     "cache_behaviour_trusted_key_groups_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TrustedKeyGroups.Enabled"),
			},
			{
				Name:     "cache_behaviour_trusted_key_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TrustedKeyGroups.Items"),
			},
			{
				Name:     "cache_behaviour_trusted_signers_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TrustedSigners.Enabled"),
			},
			{
				Name:     "cache_behaviour_trusted_signers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("DefaultCacheBehavior.TrustedSigners.Items"),
			},
			// DefaultCacheBehavior end
			{
				Name: "domain_name",
				Type: schema.TypeString,
			},
			{
				Name: "enabled",
				Type: schema.TypeBool,
			},
			{
				Name: "http_version",
				Type: schema.TypeString,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "ip_v6_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("IsIPV6Enabled"),
			},
			{
				Name: "last_modified_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "price_class",
				Type: schema.TypeString,
			},
			// Restrictions start
			{
				Name:     "restrictions_geo_restriction_restriction_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Restrictions.GeoRestriction.RestrictionType"),
			},
			{
				Name:     "restrictions_geo_restriction_restriction_items",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Restrictions.GeoRestriction.Items"),
			},
			// Restrictions End
			{
				Name: "status",
				Type: schema.TypeString,
			},
			//ViewerCertificate start
			{
				Name:     "viewer_certificate_acm_certificate_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.ACMCertificateArn"),
			},
			{
				Name:     "viewer_certificate",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.Certificate"),
			},
			{
				Name:     "viewer_certificate_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.CertificateSource"),
			},
			{
				Name:     "viewer_certificate_cloudfront_default_certificate",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("ViewerCertificate.CloudFrontDefaultCertificate"),
			},
			{
				Name:     "viewer_certificate_iam_certificate_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.IAMCertificateId"),
			},
			{
				Name:     "viewer_certificate_minimum_protocol_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.MinimumProtocolVersion"),
			},
			{
				Name:     "viewer_certificate_ssl_support_method",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ViewerCertificate.SSLSupportMethod"),
			},
			//ViewerCertificate end
			{
				Name:     "web_acl_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("WebACLId"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_cloudfront_distribution_cache_behaviours",
				Resolver: fetchCloudfrontDistributionCacheBehaviours,
				Columns: []schema.Column{
					{
						Name:     "distribution_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "path_pattern",
						Type: schema.TypeString,
					},
					{
						Name: "target_origin_id",
						Type: schema.TypeString,
					},
					{
						Name: "viewer_protocol_policy",
						Type: schema.TypeString,
					},
					{
						Name:     "allowed_methods",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("AllowedMethods.Items"),
					},
					{
						Name:     "cached_methods",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("AllowedMethods.CachedMethods.Items"),
					},
				},
			},
			{
				Name:     "aws_cache_behaviour_lambda_function_associations",
				Resolver: fetchCloudfrontDistributionDefaultCacheBehaviourLambdaFunctionAssociations,
				Columns: []schema.Column{
					{
						Name:     "distribution_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "event_type",
						Type: schema.TypeString,
					},
					{
						Name:     "lambda_function_arn",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("LambdaFunctionARN"),
					},
					{
						Name: "include_body",
						Type: schema.TypeBool,
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_custom_error_responses",
				Resolver: fetchCloudfrontDistributionCustomErrorResponses,
				Columns: []schema.Column{
					{
						Name:     "distribution_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "error_code",
						Type: schema.TypeInt,
					},
					{
						Name:     "error_caching_min_ttl",
						Type:     schema.TypeBigInt,
						Resolver: schema.PathResolver("ErrorCachingMinTTL"),
					},
					{
						Name: "response_code",
						Type: schema.TypeString,
					},
					{
						Name: "response_page_path",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_origins",
				Resolver: fetchCloudfrontDistributionOrigins,
				Options:  schema.TableCreationOptions{PrimaryKeys: []string{"distribution_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:     "distribution_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "domain_name",
						Type: schema.TypeString,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name: "connection_attempts",
						Type: schema.TypeInt,
					},
					{
						Name: "connection_timeout",
						Type: schema.TypeInt,
					},
					{
						Name:     "custom_headers",
						Type:     schema.TypeJSON,
						Resolver: resolveCloudfrontDistributionOriginCustomHeaders,
					},
					{
						Name:     "custom_origin_config_http_port",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("CustomOriginConfig.HTTPPort"),
					},
					{
						Name:     "custom_origin_config_https_port",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("CustomOriginConfig.HTTPSPort"),
					},
					{
						Name:     "custom_origin_config_protocol_policy",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginProtocolPolicy"),
					},
					{
						Name:     "custom_origin_config_keepalive_timeout",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginKeepaliveTimeout"),
					},
					{
						Name:     "custom_origin_config_read_timeout",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginReadTimeout"),
					},
					{
						Name:     "custom_origin_config_ssl_protocols",
						Type:     schema.TypeStringArray,
						Resolver: schema.PathResolver("CustomOriginConfig.OriginSslProtocols.Items"),
					},
					{
						Name: "origin_path",
						Type: schema.TypeString,
						//Resolver: schema.PathResolver("CustomOriginConfig.OriginPath"),
					},

					{
						Name:     "origin_shield_enabled",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("OriginShield.Enabled"),
					},
					{
						Name:     "origin_shield_region",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("OriginShield.OriginShieldRegion"),
					},
					{
						Name:     "s3_origin_config_origin_access_identity",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("S3OriginConfig.OriginAccessIdentity"),
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_alias_icp_recordals",
				Resolver: fetchCloudfrontDistributionAliasICPRecordals,
				Columns: []schema.Column{
					{
						Name:     "distribution_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "cname",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("CNAME"),
					},
					{
						Name:     "icp_recordal_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("ICPRecordalStatus"),
					},
				},
			},
			{
				Name:     "aws_cloudfront_distribution_origin_groups",
				Resolver: fetchCloudfrontDistributionOriginGroups,
				Options:  schema.TableCreationOptions{PrimaryKeys: []string{"distribution_cq_id", "id"}},
				Columns: []schema.Column{
					{
						Name:     "distribution_cq_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "failover_criteria_status_codes_items",
						Type:     schema.TypeIntArray,
						Resolver: resolveFailoverCriteriaStatusCodeItems,
					},
					{
						Name:     "id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Id"),
					},
					{
						Name:     "members_origin_ids",
						Type:     schema.TypeStringArray,
						Resolver: resolveCloudfrontDistributionOriginGroupMembers,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchCloudfrontDistributions(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	var config cloudfront.ListDistributionsInput
	c := meta.(*client.Client)
	svc := c.Services().Cloudfront
	for {
		response, err := svc.ListDistributions(ctx, nil, func(options *cloudfront.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		if response.DistributionList != nil {
			res <- response.DistributionList.Items
		}

		if aws.ToString(response.DistributionList.Marker) == "" {
			break
		}
		config.Marker = response.DistributionList.Marker
	}
	return nil
}

func fetchCloudfrontDistributionCacheBehaviours(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.CacheBehaviors != nil {
		res <- distribution.CacheBehaviors.Items
	}
	return nil
}

func fetchCloudfrontDistributionCustomErrorResponses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.CustomErrorResponses != nil {
		res <- distribution.CustomErrorResponses.Items
	}
	return nil
}

func resolveCloudfrontDistributionOriginCustomHeaders(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Origin)
	if r.CustomHeaders == nil {
		return nil
	}
	tags := map[string]string{}
	for _, t := range r.CustomHeaders.Items {
		tags[*t.HeaderName] = *t.HeaderValue
	}
	return resource.Set("custom_headers", tags)
}

func fetchCloudfrontDistributionOrigins(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.Origins != nil {
		res <- distribution.Origins.Items
	}
	return nil
}

func fetchCloudfrontDistributionAliasICPRecordals(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	res <- distribution.AliasICPRecordals
	return nil
}

func fetchCloudfrontDistributionOriginGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.OriginGroups != nil {
		res <- distribution.OriginGroups.Items
	}
	return nil
}

func resolveCloudfrontDistributionOriginGroupMembers(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.OriginGroup)
	if r.Members == nil {
		return nil
	}
	members := make([]string, 0, *r.Members.Quantity)
	for _, t := range r.Members.Items {
		members = append(members, *t.OriginId)
	}
	return resource.Set("members_origin_ids", members)
}

func resolveFailoverCriteriaStatusCodeItems(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.OriginGroup)
	if r.FailoverCriteria == nil || r.FailoverCriteria.StatusCodes == nil {
		return nil
	}
	members := make([]int, 0, *r.Members.Quantity)
	for _, item := range r.FailoverCriteria.StatusCodes.Items {
		members = append(members, int(item))
	}
	return resource.Set(c.Name, members)
}

func fetchCloudfrontDistributionDefaultCacheBehaviourLambdaFunctionAssociations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	distribution, ok := parent.Item.(types.DistributionSummary)
	if !ok {
		return fmt.Errorf("not cloudfront distribution")
	}
	if distribution.DefaultCacheBehavior != nil && distribution.DefaultCacheBehavior.LambdaFunctionAssociations != nil {
		res <- distribution.DefaultCacheBehavior.LambdaFunctionAssociations.Items
	}
	return nil
}
