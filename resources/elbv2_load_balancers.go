package resources

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	elbv2 "github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2"
	"github.com/aws/aws-sdk-go-v2/service/elasticloadbalancingv2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Elbv2LoadBalancers() *schema.Table {
	return &schema.Table{
		Name:         "aws_elbv2_load_balancers",
		Description:  "Information about a load balancer.",
		Resolver:     fetchElbv2LoadBalancers,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
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
				Name:        "customer_owned_ipv4_pool",
				Description: "[Application Load Balancers on Outposts] The ID of the customer-owned address pool.",
				Type:        schema.TypeString,
			},
			{
				Name:        "dns_name",
				Description: "The public DNS name of the load balancer.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("DNSName"),
			},
			{
				Name:        "ip_address_type",
				Description: "The type of IP addresses used by the subnets for your load balancer.",
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
				Description: "The nodes of an Internet-facing load balancer have public IP addresses.",
				Type:        schema.TypeString,
			},
			{
				Name:        "security_groups",
				Description: "The IDs of the security groups for the load balancer.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "state_code",
				Description: "The state code.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State.Code"),
			},
			{
				Name:        "state_reason",
				Description: "A description of the state.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State.Reason"),
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
			{
				Name:        "aws_elbv2_load_balancer_availability_zones",
				Description: "Information about an Availability Zone.",
				Resolver:    fetchElbv2LoadBalancerAvailabilityZones,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"load_balancer_cq_id", "zone_name"}},
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
						Name:        "outpost_id",
						Description: "[Application Load Balancers on Outposts] The ID of the Outpost.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_id",
						Description: "The ID of the subnet.",
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
						Name:        "aws_elbv2_load_balancer_availability_zone_addresses",
						Description: "Information about a static IP address for a load balancer.",
						Resolver:    fetchElbv2LoadBalancerAvailabilityZoneAddresses,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"load_balancer_availability_zone_cq_id", "ip_address"}},
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
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchElbv2LoadBalancers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	var config elbv2.DescribeLoadBalancersInput
	c := meta.(*client.Client)
	svc := c.Services().ELBv2
	for {
		response, err := svc.DescribeLoadBalancers(ctx, &config, func(options *elbv2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		res <- response.LoadBalancers
		if aws.ToString(response.NextMarker) == "" {
			break
		}
		config.Marker = response.NextMarker
	}
	return nil
}
func fetchElbv2LoadBalancerAvailabilityZones(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(types.LoadBalancer)
	res <- p.AvailabilityZones
	return nil
}
func fetchElbv2LoadBalancerAvailabilityZoneAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	p := parent.Item.(types.AvailabilityZone)
	res <- p.LoadBalancerAddresses
	return nil
}
