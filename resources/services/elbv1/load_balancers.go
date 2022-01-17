package elbv1

import (
	"context"
	"errors"

	"github.com/aws/smithy-go"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Elbv1LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:         "aws_elbv1_load_balancers",
		Description:  "Information about a load balancer.",
		Resolver:     fetchElbv1LoadBalancers,
		Multiplex:    client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "region", "name"}},
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
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) for the resource.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.ElasticLoadBalancingService, func(resource *schema.Resource) ([]string, error) {
					return []string{"loadbalancer", *resource.Item.(ELBv1LoadBalancerWrapper).LoadBalancerName}, nil
				}),
			},
			{
				Name:     "attributes_access_log_enabled",
				Type:     schema.TypeBool,
				Resolver: resolveElbv1loadBalancerAttributesAccessLogEnabled,
			},
			{
				Name:     "attributes_access_log_s3_bucket_name",
				Type:     schema.TypeString,
				Resolver: resolveElbv1loadBalancerAttributesAccessLogS3BucketName,
			},
			{
				Name:     "attributes_access_log_s3_bucket_prefix",
				Type:     schema.TypeString,
				Resolver: resolveElbv1loadBalancerAttributesAccessLogS3BucketPrefix,
			},
			{
				Name:     "attributes_access_log_emit_interval",
				Type:     schema.TypeInt,
				Resolver: resolveElbv1loadBalancerAttributesAccessLogEmitInterval,
			},
			{
				Name:     "attributes_connection_settings_idle_timeout",
				Type:     schema.TypeInt,
				Resolver: resolveElbv1loadBalancerAttributesConnectionSettingsIdleTimeout,
			},
			{
				Name:     "attributes_cross_zone_load_balancing_enabled",
				Type:     schema.TypeBool,
				Resolver: resolveElbv1loadBalancerAttributesCrossZoneLoadBalancingEnabled,
			},
			{
				Name:     "attributes_connection_draining_enabled",
				Type:     schema.TypeBool,
				Resolver: resolveElbv1loadBalancerAttributesConnectionDrainingEnabled,
			},
			{
				Name:     "attributes_connection_draining_timeout",
				Type:     schema.TypeInt,
				Resolver: resolveElbv1loadBalancerAttributesConnectionDrainingTimeout,
			},
			{
				Name:     "attributes_additional_attributes",
				Type:     schema.TypeJSON,
				Resolver: resolveElbv1loadBalancerAttributesAdditionalAttributes,
			},
			{
				Name: "tags",
				Type: schema.TypeJSON,
			},
			{
				Name:        "availability_zones",
				Description: "The Availability Zones for the load balancer.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "canonical_hosted_zone_name",
				Description: "The DNS name of the load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "canonical_hosted_zone_name_id",
				Description: "The ID of the Amazon Route 53 hosted zone for the load balancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CanonicalHostedZoneNameID"),
			},
			{
				Name:        "created_time",
				Description: "The date and time the load balancer was created.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "dns_name",
				Description: "The DNS name of the load balancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DNSName"),
			},
			{
				Name:        "health_check_healthy_threshold",
				Description: "The number of consecutive health checks successes required before moving the instance to the Healthy state.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HealthCheck.HealthyThreshold"),
			},
			{
				Name:        "health_check_interval",
				Description: "The approximate interval, in seconds, between health checks of an individual instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HealthCheck.Interval"),
			},
			{
				Name:        "health_check_target",
				Description: "The instance being checked.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HealthCheck.Target"),
			},
			{
				Name:        "health_check_timeout",
				Description: "The amount of time, in seconds, during which no response means a failed health check.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HealthCheck.Timeout"),
			},
			{
				Name:        "health_check_unhealthy_threshold",
				Description: "The number of consecutive health check failures required before moving the instance to the Unhealthy state.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HealthCheck.UnhealthyThreshold"),
			},
			{
				Name:        "instances",
				Description: "The IDs of the instances for the load balancer.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveElbv1loadBalancerInstances,
			},
			{
				Name:        "name",
				Description: "The name of the load balancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("LoadBalancerName"),
			},
			{
				Name:        "other_policies",
				Description: "The policies other than the stickiness policies.",
				Type:        schema.TypeStringArray,
				Resolver:    schema.PathResolver("Policies.OtherPolicies"),
			},
			{
				Name:        "scheme",
				Description: "The type of load balancer.",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_groups",
				Description: "The security groups for the load balancer.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "source_security_group_name",
				Description: "The name of the security group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSecurityGroup.GroupName"),
			},
			{
				Name:        "source_security_group_owner_alias",
				Description: "The owner of the security group.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("SourceSecurityGroup.OwnerAlias"),
			},
			{
				Name:        "subnets",
				Description: "The IDs of the subnets for the load balancer.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC for the load balancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("VPCId"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_elbv1_load_balancer_backend_server_descriptions",
				Description: "Information about the configuration of an EC2 instance.",
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"load_balancer_cq_id", "instance_port"}},
				Resolver:    fetchElbv1LoadBalancerBackendServerDescriptions,
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv1_load_balancers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of the load balancer.",
						Type:        schema.TypeString,
						Resolver:    schema.ParentResourceFieldResolver("name"),
					},
					{
						Name:        "instance_port",
						Description: "The port on which the EC2 instance is listening.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "policy_names",
						Description: "The names of the policies enabled for the EC2 instance.",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "aws_elbv1_load_balancer_listeners",
				Description: "The policies enabled for a listener.",
				Resolver:    fetchElbv1LoadBalancerListeners,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"load_balancer_cq_id", "listener_instance_port", "listener_load_balancer_port"}},
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv1_load_balancers table (FK)",
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
						Name:        "listener_instance_port",
						Description: "The port on which the instance is listening.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Listener.InstancePort"),
					},
					{
						Name:        "listener_load_balancer_port",
						Description: "The port on which the load balancer is listening.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Listener.LoadBalancerPort"),
					},
					{
						Name:        "listener_protocol",
						Description: "The load balancer transport protocol to use for routing: HTTP, HTTPS, TCP, or SSL.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Listener.Protocol"),
					},
					{
						Name:        "listener_instance_protocol",
						Description: "The protocol to use for routing traffic to instances: HTTP, HTTPS, TCP, or SSL.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Listener.InstanceProtocol"),
					},
					{
						Name:        "listener_ssl_certificate_id",
						Description: "The Amazon Resource Name (ARN) of the server certificate.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Listener.SSLCertificateId"),
					},
					{
						Name:        "policy_names",
						Description: "The policies.",
						Type:        schema.TypeStringArray,
					},
				},
			},
			{
				Name:        "aws_elbv1_load_balancer_policies_app_cookie_stickiness",
				Description: "Information about a policy for application-controlled session stickiness.",
				Resolver:    fetchElbv1LoadBalancerPoliciesAppCookieStickinessPolicies,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"load_balancer_cq_id", "cookie_name", "policy_name"}},
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv1_load_balancers table (FK)",
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
						Name:        "cookie_name",
						Description: "The name of the application cookie used for stickiness.",
						Type:        schema.TypeString,
					},
					{
						Name:        "policy_name",
						Description: "The mnemonic name for the policy being created.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_elbv1_load_balancer_policies_lb_cookie_stickiness",
				Description: "Information about a policy for duration-based session stickiness.",
				Resolver:    fetchElbv1LoadBalancerPoliciesLbCookieStickinessPolicies,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"load_balancer_cq_id", "policy_name"}},
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv1_load_balancers table (FK)",
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
						Name:        "cookie_expiration_period",
						Description: "The time period, in seconds, after which the cookie should be considered stale.",
						Type:        schema.TypeBigInt,
					},
					{
						Name:        "policy_name",
						Description: "The name of the policy.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_elbv1_load_balancer_policies",
				Description: "Information about a policy.",
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"load_balancer_cq_id", "policy_name"}},
				Resolver:    fetchElbv1LoadBalancerPolicies,
				Columns: []schema.Column{
					{
						Name:        "load_balancer_cq_id",
						Description: "Unique CloudQuery ID of aws_elbv1_load_balancers table (FK)",
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
						Name:        "policy_attribute_descriptions",
						Description: "The policy attributes.",
						Type:        schema.TypeJSON,
						Resolver:    resolveElbv1loadBalancerPolicyPolicyAttributeDescriptions,
					},
					{
						Name:        "policy_name",
						Description: "The name of the policy.",
						Type:        schema.TypeString,
					},
					{
						Name:        "policy_type_name",
						Description: "The name of the policy type.",
						Type:        schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElbv1LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().ELBv1
	processLoadBalancers := func(loadBalancers []types.LoadBalancerDescription) error {
		tagsCfg := &elbv1.DescribeTagsInput{LoadBalancerNames: make([]string, 0, len(loadBalancers))}
		for _, lb := range loadBalancers {
			tagsCfg.LoadBalancerNames = append(tagsCfg.LoadBalancerNames, *lb.LoadBalancerName)
		}
		tagsResponse, err := svc.DescribeTags(ctx, tagsCfg)
		if err != nil {
			return err
		}
		for _, lb := range loadBalancers {

			tags := getTagsByLoadBalancerName(*lb.LoadBalancerName, tagsResponse.TagDescriptions)

			loadBalancerAttributes, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv1.DescribeLoadBalancerAttributesInput{LoadBalancerName: lb.LoadBalancerName})
			if err != nil {
				var ae smithy.APIError
				// If we received any error other than LoadBalancerNotFound, we return and error
				if errors.As(err, &ae) && ae.ErrorCode() == "LoadBalancerNotFound" {
					c.Logger().Warn("elbv1 load balancer not found", "account_id", c.AccountID, "region", c.Region, "name", lb.LoadBalancerName)
					continue
				}
			}

			wrapper := ELBv1LoadBalancerWrapper{
				LoadBalancerDescription: lb,
				Tags:                    make(map[string]interface{}, len(tags)),
				Attributes:              loadBalancerAttributes.LoadBalancerAttributes,
			}

			for _, t := range tags {
				wrapper.Tags[*t.Key] = t.Value
			}
			res <- wrapper
		}
		return nil
	}

	var config elbv1.DescribeLoadBalancersInput
	for {
		response, err := svc.DescribeLoadBalancers(ctx, &config, func(options *elbv1.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}

		for i := 0; i < len(response.LoadBalancerDescriptions); i += 20 {
			end := i + 20

			if end > len(response.LoadBalancerDescriptions) {
				end = len(response.LoadBalancerDescriptions)
			}
			loadBalancers := response.LoadBalancerDescriptions[i:end]
			if err := processLoadBalancers(loadBalancers); err != nil {
				return err
			}
		}

		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}

	return nil
}
func resolveElbv1loadBalancerAttributesAccessLogEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil && r.Attributes.AccessLog == nil {
		return nil
	}
	return resource.Set(c.Name, r.Attributes.AccessLog.Enabled)
}
func resolveElbv1loadBalancerAttributesAccessLogS3BucketName(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil && r.Attributes.AccessLog == nil {
		return nil
	}
	return resource.Set(c.Name, r.Attributes.AccessLog.S3BucketName)
}
func resolveElbv1loadBalancerAttributesAccessLogS3BucketPrefix(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil && r.Attributes.AccessLog == nil {
		return nil
	}
	return resource.Set(c.Name, r.Attributes.AccessLog.S3BucketPrefix)
}
func resolveElbv1loadBalancerAttributesAccessLogEmitInterval(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil && r.Attributes.AccessLog == nil {
		return nil
	}
	return resource.Set(c.Name, r.Attributes.AccessLog.EmitInterval)
}
func resolveElbv1loadBalancerAttributesConnectionSettingsIdleTimeout(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil && r.Attributes.ConnectionSettings == nil {
		return nil
	}

	return resource.Set(c.Name, r.Attributes.ConnectionSettings.IdleTimeout)
}
func resolveElbv1loadBalancerAttributesCrossZoneLoadBalancingEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil && r.Attributes.CrossZoneLoadBalancing == nil {
		return nil
	}
	return resource.Set(c.Name, r.Attributes.CrossZoneLoadBalancing.Enabled)
}
func resolveElbv1loadBalancerAttributesConnectionDrainingEnabled(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil && r.Attributes.ConnectionDraining == nil {
		return nil
	}
	return resource.Set(c.Name, r.Attributes.ConnectionDraining.Enabled)
}
func resolveElbv1loadBalancerAttributesConnectionDrainingTimeout(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil && r.Attributes.ConnectionDraining == nil {
		return nil
	}
	return resource.Set(c.Name, r.Attributes.ConnectionDraining.Timeout)
}
func resolveElbv1loadBalancerAttributesAdditionalAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	if r.Attributes == nil {
		return nil
	}

	response := make(map[string]interface{}, len(r.Attributes.AdditionalAttributes))
	for _, a := range r.Attributes.AdditionalAttributes {
		response[*a.Key] = a.Value
	}
	return resource.Set(c.Name, response)
}
func resolveElbv1loadBalancerInstances(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	response := make([]string, 0, len(r.Instances))
	for _, i := range r.Instances {
		response = append(response, *i.InstanceId)
	}
	return resource.Set(c.Name, response)
}
func fetchElbv1LoadBalancerBackendServerDescriptions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	res <- r.BackendServerDescriptions
	return nil
}
func fetchElbv1LoadBalancerListeners(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	res <- r.ListenerDescriptions
	return nil
}
func fetchElbv1LoadBalancerPoliciesAppCookieStickinessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}

	if r.Policies == nil {
		return nil
	}
	res <- r.Policies.AppCookieStickinessPolicies
	return nil
}
func fetchElbv1LoadBalancerPoliciesLbCookieStickinessPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}

	if r.Policies == nil {
		return nil
	}
	res <- r.Policies.LBCookieStickinessPolicies
	return nil
}
func fetchElbv1LoadBalancerPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r, ok := parent.Item.(ELBv1LoadBalancerWrapper)
	if !ok {
		return errors.New("not load balancer")
	}
	c := meta.(*client.Client)
	svc := c.Services().ELBv1
	response, err := svc.DescribeLoadBalancerPolicies(ctx, &elbv1.DescribeLoadBalancerPoliciesInput{LoadBalancerName: r.LoadBalancerName}, func(options *elbv1.Options) {
		options.Region = c.Region
	})
	if err != nil {
		return err
	}
	res <- response.PolicyDescriptions
	return nil
}
func resolveElbv1loadBalancerPolicyPolicyAttributeDescriptions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r, ok := resource.Item.(types.PolicyDescription)
	if !ok {
		return errors.New("not policy description")
	}

	response := make(map[string]interface{}, len(r.PolicyAttributeDescriptions))
	for _, a := range r.PolicyAttributeDescriptions {
		response[*a.AttributeName] = a.AttributeValue
	}
	return resource.Set(c.Name, response)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

type ELBv1LoadBalancerWrapper struct {
	types.LoadBalancerDescription
	Tags       map[string]interface{}
	Attributes *types.LoadBalancerAttributes
}

func getTagsByLoadBalancerName(id string, tagsResponse []types.TagDescription) []types.Tag {
	for _, t := range tagsResponse {
		if id == *t.LoadBalancerName {
			return t.Tags
		}
	}
	return nil
}
