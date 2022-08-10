package elbv2

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/aws/aws-sdk-go-v2/service/wafv2"
	wafv2types "github.com/aws/aws-sdk-go-v2/service/wafv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/mitchellh/mapstructure"
)

type lbAttributes struct {
	AccessLogsS3Enabled                   bool   `mapstructure:"access_logs.s3.enabled"`
	AccessLogsS3Bucket                    string `mapstructure:"access_logs.s3.bucket"`
	AccessLogsS3Prefix                    string `mapstructure:"access_logs.s3.prefix"`
	DeletionProtection                    bool   `mapstructure:"deletion_protection.enabled"`
	IdleTimeout                           int32  `mapstructure:"idle_timeout.timeout_seconds"`
	RoutingHTTPDesyncMitigationMode       string `mapstructure:"routing.http.desync_mitigation_mode"`
	RoutingHTTPDropInvalidHeaderFields    bool   `mapstructure:"routing.http.drop_invalid_header_fields.enabled"`
	RoutingHTTPXAmznTLSVersionCipherSuite bool   `mapstructure:"routing.http.x_amzn_tls_version_and_cipher_suite.enabled"`
	RoutingHTTPXFFClientPort              bool   `mapstructure:"routing.http.xff_client_port.enabled"`
	RoutingHTTP2                          bool   `mapstructure:"routing.http2.enabled"`
	WAFFailOpen                           bool   `mapstructure:"waf.fail_open.enabled"`
	LoadBalancingCrossZone                bool   `mapstructure:"load_balancing.cross_zone.enabled"`
}

func Elbv2LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:         "aws_elbv2_load_balancers",
		Description:  "Information about a load balancer.",
		Resolver:     fetchElbv2LoadBalancers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		IgnoreError:  client.IgnoreCommonErrors,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Name:          "web_acl_arn",
				Description:   "arn of associated web acl",
				Type:          schema.TypeString,
				Resolver:      resolveElbv2loadBalancerWebACLArn,
				IgnoreInTests: true,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv2loadBalancerTags,
			},
			{
				Name:        "canonical_hosted_zone_id",
				Description: "The ID of the Amazon Route 53 hosted zone associated with the load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_time",
				Description: "The date and time the load balancer was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "customer_owned_ipv4_pool",
				Description:   "[Application Load Balancers on Outposts] The ID of the customer-owned address pool.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "dns_name",
				Description: "The public DNS name of the load balancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DNSName"),
			},
			{
				Name:        "ip_address_type",
				Description: "The type of IP addresses used by the subnets for your load balancer",
				Type:        schema.TypeString,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the load balancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LoadBalancerArn"),
			},
			{
				Name:        "name",
				Description: "The name of the load balancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LoadBalancerName"),
			},
			{
				Name:        "scheme",
				Description: "The nodes of an Internet-facing load balancer have public IP addresses",
				Type:        schema.TypeString,
			},
			{
				Name:          "security_groups",
				Description:   "The IDs of the security groups for the load balancer.",
				Type:          schema.TypeStringArray,
				IgnoreInTests: true,
			},
			{
				Name:        "state_code",
				Description: "The state code",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State.Code"),
			},
			{
				Name:          "state_reason",
				Description:   "A description of the state.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("State.Reason"),
				IgnoreInTests: true,
			},
			{
				Name:        "type",
				Description: "The type of load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the load balancer.",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			Elbv2Listeners(),
			{
				Name:        "aws_elbv2_load_balancer_availability_zones",
				Description: "Information about an Availability Zone.",
				Resolver:    schema.PathTableResolver("AvailabilityZones"),
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv2_load_balancers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "load_balance_name",
						Description: "The name of the load balancer.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("name"),
					},
					{
						Name:          "outpost_id",
						Description:   "[Application Load Balancers on Outposts] The ID of the Outpost.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "subnet_id",
						Description: "The ID of the subnet",
						Type:        schema.TypeString,
					},
					{
						Name:        "zone_name",
						Description: "The name of the Availability Zone.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_elbv2_load_balancer_availability_zone_addresses",
						Description:   "Information about a static IP address for a load balancer.",
						Resolver:      schema.PathTableResolver("LoadBalancerAddresses"),
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "load_balancer_availability_zone_cq_id",
								Description: "Unique CloudQuery ID of aws_elbv2_load_balancer_availability_zones table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "zone_name",
								Description: "The name of the Availability Zone..",
								Type:        schema.TypeString,
								Resolver:    schema.ParentResourceFieldResolver("zone_name"),
							},
							{
								Name:        "allocation_id",
								Description: "[Network Load Balancers] The allocation ID of the Elastic IP address for an internal-facing load balancer.",
								Type:        schema.TypeString,
							},
							{
								Name:        "ipv6_address",
								Description: "[Network Load Balancers] The IPv6 address.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("IPv6Address"),
							},
							{
								Name:        "ip_address",
								Description: "The static IP address.",
								Type:        schema.TypeString,
							},
							{
								Name:        "private_ipv4_address",
								Description: "[Network Load Balancers] The private IPv4 address for an internal load balancer.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("PrivateIPv4Address"),
							},
						},
					},
				},
			},
			{
				Name:        "aws_elbv2_load_balancer_attributes",
				Description: "Load balancer attributes",
				Resolver:    fetchElbv2LoadBalancerAttributes,
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv2_load_balancers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "access_logs_s3_enabled",
						Description: "Indicates whether access logs stored in Amazon S3 are enabled.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "access_logs_s3_bucket",
						Description: "The name of the Amazon S3 bucket for the access logs.",
						Type:        schema.TypeString,
					},
					{
						Name:        "access_logs_s3_prefix",
						Description: "The prefix for the location in the Amazon S3 bucket.",
						Type:        schema.TypeString,
					},
					{
						Name:        "deletion_protection",
						Description: "Indicates whether deletion protection is enabled.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "idle_timeout",
						Description: "The idle timeout value, in seconds.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "routing_http_desync_mitigation_mode",
						Description: "Determines how the load balancer handles requests that might pose a security risk to your application.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RoutingHTTPDesyncMitigationMode"),
					},
					{
						Name:        "routing_http_drop_invalid_header_fields",
						Description: "Indicates whether HTTP headers with header fields that are not valid are removed by the load balancer.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("RoutingHTTPDropInvalidHeaderFields"),
					},
					{
						Name:        "routing_http_xamzntls_enabled",
						Description: "Indicates whether the two headers (x-amzn-tls-{version,cipher-suite}) are added to the client request before sending it to the target.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("RoutingHTTPXAmznTLSVersionCipherSuite"),
					},
					{
						Name:        "routing_http_xff_client_port",
						Description: "Indicates whether the X-Forwarded-For header should preserve the source port that the client used to connect to the load balancer.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("RoutingHTTPXFFClientPort"),
					},
					{
						Name:        "routing_http2",
						Description: "Indicates whether HTTP/2 is enabled.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("RoutingHTTP2"),
					},
					{
						Name:        "waf_fail_open",
						Description: "Indicates whether to allow a AWS WAF-enabled load balancer to route requests to targets if it is unable to forward the request to AWS WAF.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("WAFFailOpen"),
					},
					{
						Name:        "load_balancing_cross_zone",
						Description: "Indicates whether cross-zone load balancing is enabled.",
						Type:        schema.TypeBool,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElbv2LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config elbv2.DescribeLoadBalancersInput
	c := meta.(*client.Client)
	svc := c.Services().ELBv2
	for {
		response, err := svc.DescribeLoadBalancers(ctx, &config, func(options *elbv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- response.LoadBalancers
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
func resolveElbv2loadBalancerWebACLArn(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	p := resource.Item.(types.LoadBalancer)
	// only application load balancer can have web acl arn
	if p.Type != types.LoadBalancerTypeEnumApplication {
		return nil
	}
	cl := meta.(*client.Client).Services().WafV2
	input := wafv2.GetWebACLForResourceInput{ResourceArn: p.LoadBalancerArn}
	response, err := cl.GetWebACLForResource(ctx, &input, func(options *wafv2.Options) {})
	if err != nil {
		var exc *wafv2types.WAFNonexistentItemException
		if errors.As(err, &exc) {
			if exc.ErrorCode() == "WAFNonexistentItemException" {
				return nil
			}
		}

		return diag.WrapError(err)
	}
	if response.WebACL == nil {
		return nil
	}

	return diag.WrapError(resource.Set(c.Name, response.WebACL.ARN))
}
func resolveElbv2loadBalancerTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	cl := meta.(*client.Client)
	region := cl.Region
	svc := cl.Services().ELBv2
	loadBalancer := resource.Item.(types.LoadBalancer)
	tagsOutput, err := svc.DescribeTags(ctx, &elbv2.DescribeTagsInput{
		ResourceArns: []string{
			*loadBalancer.LoadBalancerArn,
		},
	}, func(o *elbv2.Options) {
		o.Region = region
	})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return diag.WrapError(err)
	}
	if len(tagsOutput.TagDescriptions) == 0 {
		return nil
	}
	tags := make(map[string]*string)
	for _, td := range tagsOutput.TagDescriptions {
		for _, s := range td.Tags {
			tags[*s.Key] = s.Value
		}
	}

	return diag.WrapError(resource.Set(c.Name, tags))
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func fetchElbv2LoadBalancerAttributes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	lb := parent.Item.(types.LoadBalancer)
	c := meta.(*client.Client)
	svc := c.Services().ELBv2
	result, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv2.DescribeLoadBalancerAttributesInput{LoadBalancerArn: lb.LoadBalancerArn}, func(options *elbv2.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	m := make(map[string]interface{})
	for _, a := range result.Attributes {
		m[*a.Key] = *a.Value
	}
	var attrs lbAttributes
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{WeaklyTypedInput: true, Result: &attrs})
	if err != nil {
		return diag.WrapError(err)
	}
	if err := dec.Decode(m); err != nil {
		return diag.WrapError(err)
	}
	res <- attrs
	return nil
}
