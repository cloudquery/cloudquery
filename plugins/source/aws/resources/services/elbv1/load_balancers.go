package elbv1

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv1 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancing/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

type ELBv1LoadBalancerWrapper struct {
	types.LoadBalancerDescription
	Tags       map[string]string
	Attributes *types.LoadBalancerAttributes
}

func Elbv1LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:          "aws_elbv1_load_balancers",
		Description:   "Information about a load balancer.",
		Resolver:      fetchElbv1LoadBalancers,
		Multiplex:     client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
		IgnoreInTests: true,
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:            "region",
				Description:     "The AWS Region of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSRegion,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Name:     "attributes",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Attributes"),
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
				Name:     "health_check",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("HealthCheck"),
			},
			{
				Name:     "instances",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Instances"),
			},
			{
				Name:            "name",
				Description:     "The name of the load balancer.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("LoadBalancerName"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:     "policies",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Policies"),
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
				Name:     "source_security_group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceSecurityGroup"),
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
			{
				Name:        "backend_server_descriptions",
				Description: "Information about the configuration of an EC2 instance.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("BackendServerDescriptions"),
			},
			{
				Name:        "listener_descriptions",
				Description: "The policies enabled for a listener.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ListenerDescriptions"),
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_elbv1_load_balancer_policies",
				Description:   "Information about a policy.",
				Resolver:      fetchElbv1LoadBalancerPolicies,
				IgnoreInTests: true,
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
//
//	Table Resolver Functions
//
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
			loadBalancerAttributes, err := svc.DescribeLoadBalancerAttributes(ctx, &elbv1.DescribeLoadBalancerAttributesInput{LoadBalancerName: lb.LoadBalancerName})
			if err != nil {
				if c.IsNotFoundError(err) {
					continue
				}
				return err
			}

			wrapper := ELBv1LoadBalancerWrapper{
				LoadBalancerDescription: lb,
				Tags:                    client.TagsToMap(getTagsByLoadBalancerName(*lb.LoadBalancerName, tagsResponse.TagDescriptions)),
				Attributes:              loadBalancerAttributes.LoadBalancerAttributes,
			}

			res <- wrapper
		}
		return nil
	}

	var config elbv1.DescribeLoadBalancersInput
	for {
		response, err := svc.DescribeLoadBalancers(ctx, &config)
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

func resolveElbv1loadBalancerAttributesAdditionalAttributes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(ELBv1LoadBalancerWrapper)
	if r.Attributes == nil {
		return nil
	}

	response := make(map[string]interface{}, len(r.Attributes.AdditionalAttributes))
	for _, a := range r.Attributes.AdditionalAttributes {
		response[*a.Key] = a.Value
	}
	return resource.Set(c.Name, response)
}

func fetchElbv1LoadBalancerPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(ELBv1LoadBalancerWrapper)
	c := meta.(*client.Client)
	svc := c.Services().ELBv1
	response, err := svc.DescribeLoadBalancerPolicies(ctx, &elbv1.DescribeLoadBalancerPoliciesInput{LoadBalancerName: r.LoadBalancerName})
	if err != nil {
		return err
	}
	res <- response.PolicyDescriptions
	return nil
}
func resolveElbv1loadBalancerPolicyPolicyAttributeDescriptions(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.PolicyDescription)

	response := make(map[string]interface{}, len(r.PolicyAttributeDescriptions))
	for _, a := range r.PolicyAttributeDescriptions {
		response[*a.AttributeName] = a.AttributeValue
	}
	return resource.Set(c.Name, response)
}

// ====================================================================================================================
//                                                  User Defined Helpers
// ====================================================================================================================

func getTagsByLoadBalancerName(id string, tagsResponse []types.TagDescription) []types.Tag {
	for _, t := range tagsResponse {
		if id == *t.LoadBalancerName {
			return t.Tags
		}
	}
	return nil
}
