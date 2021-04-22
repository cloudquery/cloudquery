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
		Resolver:     fetchElbv2LoadBalancers,
		Multiplex:    client.AccountRegionMultiplex,
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
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
				Name: "canonical_hosted_zone_id",
				Type: schema.TypeString,
			},
			{
				Name: "created_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name: "customer_owned_ipv4_pool",
				Type: schema.TypeString,
			},
			{
				Name:     "dns_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DNSName"),
			},
			{
				Name: "ip_address_type",
				Type: schema.TypeString,
			},
			{
				Name: "load_balancer_arn",
				Type: schema.TypeString,
			},
			{
				Name: "load_balancer_name",
				Type: schema.TypeString,
			},
			{
				Name: "scheme",
				Type: schema.TypeString,
			},
			{
				Name: "security_groups",
				Type: schema.TypeStringArray,
			},
			{
				Name:     "state_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State.Code"),
			},
			{
				Name:     "state_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State.Reason"),
			},
			{
				Name: "type",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_elbv2_load_balancer_availability_zones",
				Resolver: fetchElbv2LoadBalancerAvailabilityZones,
				Columns: []schema.Column{
					{
						Name:     "load_balancer_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "outpost_id",
						Type: schema.TypeString,
					},
					{
						Name: "subnet_id",
						Type: schema.TypeString,
					},
					{
						Name: "zone_name",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_elbv2_load_balancer_availability_zone_addresses",
						Resolver: fetchElbv2LoadBalancerAvailabilityZoneAddresses,
						Columns: []schema.Column{
							{
								Name:     "load_balancer_availability_zone_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "allocation_id",
								Type: schema.TypeString,
							},
							{
								Name:     "ip_v6_address",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("IPv6Address"),
							},
							{
								Name: "ip_address",
								Type: schema.TypeString,
							},
							{
								Name:     "private_ip_v4_address",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("PrivateIPv4Address"),
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
