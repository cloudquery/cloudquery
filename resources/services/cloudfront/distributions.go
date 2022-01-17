package cloudfront

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
		Description:  "A summary of the information about a CloudFront distribution.",
		Resolver:     fetchCloudfrontDistributions,
		Multiplex:    client.AccountMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:        "arn",
				Description: "The ARN (Amazon Resource Name) for the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("ARN"),
			},
			{
				Name:        "caller_reference",
				Description: "A unique value (for example, a date-time stamp) that ensures that the request can't be replayed",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.CallerReference"),
			},
			{
				Name:        "comment",
				Description: "Any comments you want to include about the distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Comment"),
			},
			{
				Name:        "cache_behavior_target_origin_id",
				Description: "The value of ID for the origin that you want CloudFront to route requests to when they use the default cache behavior.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TargetOriginId"),
			},
			{
				Name:        "cache_behavior_viewer_protocol_policy",
				Description: "The protocol that viewers can use to access the files in the origin specified by TargetOriginId when a request matches the path pattern in PathPattern",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ViewerProtocolPolicy"),
			},
			{
				Name:        "cache_behavior_allowed_methods",
				Description: "A complex type that contains the HTTP methods that you want CloudFront to process and forward to your origin.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.AllowedMethods.Items"),
			},
			{
				Name:        "cache_behavior_allowed_methods_cached_methods",
				Description: "A complex type that contains the HTTP methods that you want CloudFront to cache responses to.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.AllowedMethods.CachedMethods.Items"),
			},
			{
				Name:          "cache_behavior_cache_policy_id",
				Description:   "The unique identifier of the cache policy that is attached to the default cache behavior",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DistributionConfig.DefaultCacheBehavior.CachePolicyId"),
				IgnoreInTests: true,
			},
			{
				Name:        "cache_behavior_compress",
				Description: "Whether you want CloudFront to automatically compress certain files for this cache behavior",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.Compress"),
			},
			{
				Name:        "cache_behavior_default_ttl",
				Description: "This field is deprecated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.DefaultTTL"),
			},
			{
				Name:        "cache_behavior_field_level_encryption_id",
				Description: "The value of ID for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for the default cache behavior.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.FieldLevelEncryptionId"),
			},
			{
				Name:        "cache_behavior_forwarded_values_cookies_forward",
				Description: "This field is deprecated",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies.Forward"),
			},
			{
				Name:          "cache_behavior_forwarded_values_cookies_whitelisted_names",
				Description:   "A list of cookie names.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.Cookies.WhitelistedNames.Items"),
				IgnoreInTests: true,
			},
			{
				Name:        "cache_behavior_forwarded_values_query_string",
				Description: "This field is deprecated",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.QueryString"),
			},
			{
				Name:          "cache_behavior_forwarded_values_headers",
				Description:   "A list of HTTP header names.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.Headers.Items"),
				IgnoreInTests: true,
			},
			{
				Name:          "cache_behavior_forwarded_values_query_string_cache_keys",
				Description:   "A list that contains the query string parameters that you want CloudFront to use as a basis for caching for a cache behavior",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("DistributionConfig.DefaultCacheBehavior.ForwardedValues.QueryStringCacheKeys.Items"),
				IgnoreInTests: true,
			},
			{
				Name:        "cache_behavior_max_ttl",
				Description: "This field is deprecated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.MaxTTL"),
			},
			{
				Name:        "cache_behavior_min_ttl",
				Description: "This field is deprecated",
				Type:        schema.TypeBigInt,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.MinTTL"),
			},
			{
				Name:          "cache_behavior_origin_request_policy_id",
				Description:   "The unique identifier of the origin request policy that is attached to the default cache behavior",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DistributionConfig.DefaultCacheBehavior.OriginRequestPolicyId"),
				IgnoreInTests: true,
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
				Name:        "cache_behavior_trusted_key_groups_enabled",
				Description: "This field is true if any of the key groups in the list have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies. If not, this field is false.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedKeyGroups.Enabled"),
			},
			{
				Name:          "cache_behavior_trusted_key_groups",
				Description:   "A list of key groups identifiers.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedKeyGroups.Items"),
				IgnoreInTests: true,
			},
			{
				Name:        "cache_behavior_trusted_signers_enabled",
				Description: "This field is true if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedSigners.Enabled"),
			},
			{
				Name:          "cache_behavior_trusted_signers",
				Description:   "A list of AWS account identifiers.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("DistributionConfig.DefaultCacheBehavior.TrustedSigners.Items"),
				IgnoreInTests: true,
			},
			{
				Name:        "enabled",
				Description: "From this field, you can enable or disable the selected distribution.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.Enabled"),
			},
			{
				Name:          "aliases",
				Description:   "A complex type that contains the CNAME aliases, if any, that you want to associate with this distribution.",
				Type:          schema.TypeStringArray,
				Resolver:      schema.PathResolver("DistributionConfig.Aliases.Items"),
				IgnoreInTests: true,
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
				Name:        "ipv6_enabled",
				Description: "If you want CloudFront to respond to IPv6 DNS requests with an IPv6 address for your distribution, specify true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.IsIPV6Enabled"),
			},
			{
				Name:        "logging_bucket",
				Description: "The Amazon S3 bucket to store the access logs in, for example, myawslogbucket.s3.amazonaws.com.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.Bucket"),
			},
			{
				Name:        "logging_enabled",
				Description: "Specifies whether you want CloudFront to save access logs to an Amazon S3 bucket",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.Enabled"),
			},
			{
				Name:        "logging_include_cookies",
				Description: "Specifies whether you want CloudFront to include cookies in access logs, specify true for IncludeCookies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.IncludeCookies"),
			},
			{
				Name:        "logging_prefix",
				Description: "An optional string that you want CloudFront to prefix to the access log filenames for this distribution, for example, myprefix/",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.Logging.Prefix"),
			},
			{
				Name:        "price_class",
				Description: "The price class that corresponds with the maximum price that you want to pay for CloudFront service",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.PriceClass"),
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
				Name:          "viewer_certificate",
				Description:   "This field is deprecated",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("DistributionConfig.ViewerCertificate.Certificate"),
				IgnoreInTests: true,
			},
			{
				Name:        "viewer_certificate_source",
				Description: "This field is deprecated",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.CertificateSource"),
			},
			{
				Name:        "viewer_certificate_cloudfront_default_certificate",
				Description: "If the distribution uses the CloudFront domain name such as d111111abcdef8.cloudfront.net, set this field to true",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("DistributionConfig.ViewerCertificate.CloudFrontDefaultCertificate"),
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
				Name:        "web_acl_id",
				Description: "A unique identifier that specifies the AWS WAF web ACL, if any, to associate with this distribution",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DistributionConfig.WebACLId"),
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
				Name:        "active_trusted_key_groups_enabled",
				Description: "This field is true if any of the key groups have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ActiveTrustedKeyGroups.Enabled"),
			},
			{
				Name:        "active_trusted_key_groups",
				Description: "A list of key groups, including the identifiers of the public keys in each key group that CloudFront can use to verify the signatures of signed URLs and signed cookies.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCloudfrontDistributionsActiveTrustedKeyGroups,
			},
			{
				Name:        "active_trusted_signers_enabled",
				Description: "This field is true if any of the AWS accounts in the list have active CloudFront key pairs that CloudFront can use to verify the signatures of signed URLs and signed cookies",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ActiveTrustedSigners.Enabled"),
			},
			{
				Name:        "active_trusted_signers",
				Description: "A list of AWS accounts and the identifiers of active CloudFront key pairs in each account that CloudFront can use to verify the signatures of signed URLs and signed cookies.",
				Type:        schema.TypeJSON,
				Resolver:    resolveCloudfrontDistributionsActiveTrustedSigners,
			},
			{
				Name:        "alias_icp_recordals",
				Description: "AWS services in China customers must file for an Internet Content Provider (ICP) recordal if they want to serve content publicly on an alternate domain name, also known as a CNAME, that they've added to CloudFront",
				Type:        schema.TypeJSON,
				Resolver:    resolveCloudfrontDistributionsAliasIcpRecordals,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_cloudfront_distribution_default_cache_behavior_functions",
				Description:   "A complex type that contains a Lambda function association.",
				Resolver:      fetchCloudfrontDistributionDefaultCacheBehaviorLambdaFunctions,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "event_type",
						Description: "Specifies the event type that triggers a Lambda function invocation",
						Type:        schema.TypeString,
					},
					{
						Name:        "lambda_function_arn",
						Description: "The ARN of the Lambda function",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("LambdaFunctionARN"),
					},
					{
						Name:        "include_body",
						Description: "A flag that allows a Lambda function to have read access to the body content. For more information, see Accessing the Request Body by Choosing the Include Body Option (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.",
						Type:        schema.TypeBool,
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_origins",
				Description: "An origin",
				Resolver:    fetchCloudfrontDistributionOrigins,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "domain_name",
						Description: "The domain name for the origin",
						Type:        schema.TypeString,
					},
					{
						Name:        "id",
						Description: "A unique identifier for the origin",
						Type:        schema.TypeString,
					},
					{
						Name:        "connection_attempts",
						Description: "The number of times that CloudFront attempts to connect to the origin",
						Type:        schema.TypeInt,
					},
					{
						Name:        "connection_timeout",
						Description: "The number of seconds that CloudFront waits when trying to establish a connection to the origin",
						Type:        schema.TypeInt,
					},
					{
						Name:        "custom_headers",
						Description: "A list of HTTP header names and values that CloudFront adds to the requests that it sends to the origin",
						Type:        schema.TypeJSON,
						Resolver:    resolveCloudfrontDistributionOriginsCustomHeaders,
					},
					{
						Name:          "custom_origin_config_http_port",
						Description:   "The HTTP port that CloudFront uses to connect to the origin",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("CustomOriginConfig.HTTPPort"),
						IgnoreInTests: true,
					},
					{
						Name:          "custom_origin_config_https_port",
						Description:   "The HTTPS port that CloudFront uses to connect to the origin",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("CustomOriginConfig.HTTPSPort"),
						IgnoreInTests: true,
					},
					{
						Name:        "custom_origin_config_protocol_policy",
						Description: "Specifies the protocol (HTTP or HTTPS) that CloudFront uses to connect to the origin",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("CustomOriginConfig.OriginProtocolPolicy"),
					},
					{
						Name:          "custom_origin_config_keepalive_timeout",
						Description:   "Specifies how long, in seconds, CloudFront persists its connection to the origin",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("CustomOriginConfig.OriginKeepaliveTimeout"),
						IgnoreInTests: true,
					},
					{
						Name:          "custom_origin_config_read_timeout",
						Description:   "Specifies how long, in seconds, CloudFront waits for a response from the origin. This is also known as the origin response timeout",
						Type:          schema.TypeInt,
						Resolver:      schema.PathResolver("CustomOriginConfig.OriginReadTimeout"),
						IgnoreInTests: true,
					},
					{
						Name:          "custom_origin_config_ssl_protocols",
						Description:   "A list that contains allowed SSL/TLS protocols for this distribution.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("CustomOriginConfig.OriginSslProtocols.Items"),
						IgnoreInTests: true,
					},
					{
						Name:        "origin_path",
						Description: "An optional path that CloudFront appends to the origin domain name when CloudFront requests content from the origin",
						Type:        schema.TypeString,
					},
					{
						Name:        "origin_shield_enabled",
						Description: "A flag that specifies whether Origin Shield is enabled",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("OriginShield.Enabled"),
					},
					{
						Name:          "origin_shield_region",
						Description:   "The AWS Region for Origin Shield",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("OriginShield.OriginShieldRegion"),
						IgnoreInTests: true,
					},
					{
						Name:        "s3_origin_config_origin_access_identity",
						Description: "The CloudFront origin access identity to associate with the origin",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("S3OriginConfig.OriginAccessIdentity"),
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_cache_behaviors",
				Description: "A complex type that describes how CloudFront processes requests",
				Resolver:    fetchCloudfrontDistributionCacheBehaviors,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "path_pattern",
						Description: "The pattern (for example, images/*.jpg) that specifies which requests to apply the behavior to",
						Type:        schema.TypeString,
					},
					{
						Name:        "target_origin_id",
						Description: "The value of ID for the origin that you want CloudFront to route requests to when they match this cache behavior.",
						Type:        schema.TypeString,
					},
					{
						Name:        "viewer_protocol_policy",
						Description: "The protocol that viewers can use to access the files in the origin specified by TargetOriginId when a request matches the path pattern in PathPattern",
						Type:        schema.TypeString,
					},
					{
						Name:        "allowed_methods",
						Description: "A complex type that contains the HTTP methods that you want CloudFront to process and forward to your origin.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("AllowedMethods.Items"),
					},
					{
						Name:        "cached_methods",
						Description: "A complex type that contains the HTTP methods that you want CloudFront to cache responses to.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("AllowedMethods.CachedMethods.Items"),
					},
					{
						Name:          "cache_policy_id",
						Description:   "The unique identifier of the cache policy that is attached to this cache behavior",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "compress",
						Description: "Whether you want CloudFront to automatically compress certain files for this cache behavior",
						Type:        schema.TypeBool,
					},
					{
						Name:        "default_ttl",
						Description: "This field is deprecated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("DefaultTTL"),
					},
					{
						Name:        "field_level_encryption_id",
						Description: "The value of ID for the field-level encryption configuration that you want CloudFront to use for encrypting specific fields of data for this cache behavior.",
						Type:        schema.TypeString,
					},
					{
						Name:        "forwarded_values_cookies_forward",
						Description: "This field is deprecated",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("ForwardedValues.Cookies.Forward"),
					},
					{
						Name:          "forwarded_values_cookies_whitelisted_names",
						Description:   "A list of cookie names.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("ForwardedValues.Cookies.WhitelistedNames.Items"),
						IgnoreInTests: true,
					},
					{
						Name:        "forwarded_values_query_string",
						Description: "This field is deprecated",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("ForwardedValues.QueryString"),
					},
					{
						Name:        "forwarded_values_headers",
						Description: "A list of HTTP header names.",
						Type:        schema.TypeStringArray,
						Resolver:    schema.PathResolver("ForwardedValues.Headers.Items"),
					},
					{
						Name:          "forwarded_values_query_string_cache_keys",
						Description:   "A list that contains the query string parameters that you want CloudFront to use as a basis for caching for a cache behavior",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("ForwardedValues.QueryStringCacheKeys.Items"),
						IgnoreInTests: true,
					},
					{
						Name:        "max_ttl",
						Description: "This field is deprecated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("MaxTTL"),
					},
					{
						Name:        "min_ttl",
						Description: "This field is deprecated",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("MinTTL"),
					},
					{
						Name:          "origin_request_policy_id",
						Description:   "The unique identifier of the origin request policy that is attached to this cache behavior",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:          "realtime_log_config_arn",
						Description:   "The Amazon Resource Name (ARN) of the real-time log configuration that is attached to this cache behavior",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "smooth_streaming",
						Description: "Indicates whether you want to distribute media files in the Microsoft Smooth Streaming format using the origin that is associated with this cache behavior. If so, specify true; if not, specify false",
						Type:        schema.TypeBool,
					},
					{
						Name:        "trusted_key_groups_enabled",
						Description: "This field is true if any of the key groups in the list have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies. If not, this field is false.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("TrustedKeyGroups.Enabled"),
					},
					{
						Name:          "trusted_key_groups",
						Description:   "A list of key groups identifiers.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("TrustedKeyGroups.Items"),
						IgnoreInTests: true,
					},
					{
						Name:        "trusted_signers_enabled",
						Description: "This field is true if any of the AWS accounts have public keys that CloudFront can use to verify the signatures of signed URLs and signed cookies",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("TrustedSigners.Enabled"),
					},
					{
						Name:          "trusted_signers",
						Description:   "A list of AWS account identifiers.",
						Type:          schema.TypeStringArray,
						Resolver:      schema.PathResolver("TrustedSigners.Items"),
						IgnoreInTests: true,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_cloudfront_distribution_cache_behavior_lambda_functions",
						Description:   "A complex type that contains a Lambda function association.",
						Resolver:      fetchCloudfrontDistributionCacheBehaviorLambdaFunctions,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "distribution_cache_behavior_cq_id",
								Description: "Unique CloudQuery ID of aws_cloudfront_distribution_cache_behaviors table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "event_type",
								Description: "Specifies the event type that triggers a Lambda function invocation",
								Type:        schema.TypeString,
							},
							{
								Name:        "lambda_function_arn",
								Description: "The ARN of the Lambda function",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("LambdaFunctionARN"),
							},
							{
								Name:        "include_body",
								Description: "A flag that allows a Lambda function to have read access to the body content. For more information, see Accessing the Request Body by Choosing the Include Body Option (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/lambda-include-body-access.html) in the Amazon CloudFront Developer Guide.",
								Type:        schema.TypeBool,
							},
						},
					},
				},
			},
			{
				Name:        "aws_cloudfront_distribution_custom_error_responses",
				Description: "A complex type that controls:  * Whether CloudFront replaces HTTP status codes in the 4xx and 5xx range with custom error messages before returning the response to the viewer.  * How long CloudFront caches HTTP status codes in the 4xx and 5xx range.  For more information about custom error pages, see Customizing Error Responses (https://docs.aws.amazon.com/AmazonCloudFront/latest/DeveloperGuide/custom-error-pages.html) in the Amazon CloudFront Developer Guide.",
				Resolver:    fetchCloudfrontDistributionCustomErrorResponses,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "error_code",
						Description: "The HTTP status code for which you want to specify a custom error page and/or a caching duration.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "error_caching_min_ttl",
						Description: "The minimum amount of time, in seconds, that you want CloudFront to cache the HTTP status code specified in ErrorCode",
						Type:        schema.TypeBigInt,
						Resolver:    schema.PathResolver("ErrorCachingMinTTL"),
					},
					{
						Name:        "response_code",
						Description: "The HTTP status code that you want CloudFront to return to the viewer along with the custom error page",
						Type:        schema.TypeString,
					},
					{
						Name:        "response_page_path",
						Description: "The path to the custom error page that you want CloudFront to return to a viewer when your origin returns the HTTP status code specified by ErrorCode, for example, /4xx-errors/403-forbidden.html",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_cloudfront_distribution_origin_groups",
				Description:   "An origin group includes two origins (a primary origin and a second origin to failover to) and a failover criteria that you specify",
				Resolver:      fetchCloudfrontDistributionOriginGroups,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "distribution_cq_id",
						Description: "Unique CloudQuery ID of aws_cloudfront_distributions table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "failover_criteria_status_codes",
						Description: "The items (status codes) for an origin group.",
						Type:        schema.TypeIntArray,
						Resolver:    resolveCloudfrontDistributionOriginGroupsFailoverCriteriaStatusCodes,
					},
					{
						Name:        "id",
						Description: "The origin group's ID.",
						Type:        schema.TypeString,
					},
					{
						Name:        "members_origin_ids",
						Description: "Items (origins) in an origin group.",
						Type:        schema.TypeStringArray,
						Resolver:    resolveCloudfrontDistributionOriginGroupsMembersOriginIds,
					},
				},
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
		response, err := svc.ListDistributions(ctx, &config, func(options *cloudfront.Options) {
			options.Region = c.Region
		})
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
	distribution, ok := resource.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("expected types.Distribution but got %T", resource.Item)
	}

	client := meta.(*client.Client)
	svc := client.Services().Cloudfront
	response, err := svc.ListTagsForResource(ctx, &cloudfront.ListTagsForResourceInput{
		Resource: distribution.ARN,
	}, func(options *cloudfront.Options) {
		options.Region = client.Region
	})
	if err != nil {
		return err
	}

	tags := make(map[string]interface{})
	for _, t := range response.Tags.Items {
		tags[*t.Key] = t.Value
	}
	return resource.Set(c.Name, tags)
}
func resolveCloudfrontDistributionsActiveTrustedKeyGroups(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	distribution, ok := resource.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("not types.Distribution")
	}
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
	distribution, ok := resource.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("not types.Distribution")
	}
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
	distribution, ok := resource.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("not types.Distribution")
	}
	j := map[string]interface{}{}
	for _, a := range distribution.AliasICPRecordals {
		j[*a.CNAME] = a.ICPRecordalStatus
	}
	return resource.Set(c.Name, j)
}
func fetchCloudfrontDistributionDefaultCacheBehaviorLambdaFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("types.Distribution")
	}
	res <- r.DistributionConfig.DefaultCacheBehavior.LambdaFunctionAssociations.Items
	return nil
}
func fetchCloudfrontDistributionOrigins(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	distribution, ok := parent.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("not types.Distribution")
	}
	if distribution.DistributionConfig.Origins == nil {
		return nil
	}
	res <- distribution.DistributionConfig.Origins.Items
	return nil
}
func resolveCloudfrontDistributionOriginsCustomHeaders(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Origin)
	if r.CustomHeaders == nil {
		return nil
	}
	tags := map[string]interface{}{}
	for _, t := range r.CustomHeaders.Items {
		tags[*t.HeaderName] = *t.HeaderValue
	}
	return resource.Set(c.Name, tags)
}
func fetchCloudfrontDistributionCacheBehaviors(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	distribution, ok := parent.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("not types.Distribution")
	}
	if distribution.DistributionConfig.CacheBehaviors != nil {
		res <- distribution.DistributionConfig.CacheBehaviors.Items
	}
	return nil
}
func fetchCloudfrontDistributionCacheBehaviorLambdaFunctions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cacheBehavior, ok := parent.Item.(types.CacheBehavior)
	if !ok {
		return fmt.Errorf("not types.CacheBehavior")
	}
	if cacheBehavior.LambdaFunctionAssociations == nil {
		return nil
	}
	res <- cacheBehavior.LambdaFunctionAssociations.Items
	return nil
}
func fetchCloudfrontDistributionCustomErrorResponses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	distribution, ok := parent.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("not types.Distribution")
	}
	if distribution.DistributionConfig.CustomErrorResponses != nil {
		res <- distribution.DistributionConfig.CustomErrorResponses.Items
	}
	return nil
}
func fetchCloudfrontDistributionOriginGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	distribution, ok := parent.Item.(types.Distribution)
	if !ok {
		return fmt.Errorf("not types.Distribution")
	}
	if distribution.DistributionConfig.OriginGroups != nil {
		res <- distribution.DistributionConfig.OriginGroups.Items
	}
	return nil
}
func resolveCloudfrontDistributionOriginGroupsFailoverCriteriaStatusCodes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	origin, ok := resource.Item.(types.OriginGroup)
	if !ok {
		return fmt.Errorf("not types.OriginGroup")
	}
	if origin.FailoverCriteria == nil || origin.FailoverCriteria.StatusCodes == nil {
		return nil
	}
	data := make([]int, 0, *origin.FailoverCriteria.StatusCodes.Quantity)
	for _, i := range origin.FailoverCriteria.StatusCodes.Items {
		data = append(data, int(i))
	}
	return resource.Set(c.Name, data)
}
func resolveCloudfrontDistributionOriginGroupsMembersOriginIds(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.OriginGroup)
	if r.Members == nil {
		return nil
	}
	members := make([]string, 0, *r.Members.Quantity)
	for _, t := range r.Members.Items {
		members = append(members, *t.OriginId)
	}
	return resource.Set(c.Name, members)
}
