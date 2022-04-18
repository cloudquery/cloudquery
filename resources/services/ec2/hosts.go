package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource hosts --config gen.hcl --output .
func Hosts() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_hosts",
		Description:   "Describes the properties of the Dedicated Host.",
		Resolver:      fetchEc2Hosts,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:   client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter:  client.DeleteAccountRegionFilter,
		IgnoreInTests: true,
		Options:       schema.TableCreationOptions{PrimaryKeys: []string{"arn"}},
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
				Description: "The Amazon Resource Name (ARN) for the dedicated host.",
				Type:        schema.TypeString,
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"dedicated-host", *resource.Item.(types.Host).HostId}, nil
				}),
			},
			{
				Name:        "allocation_time",
				Description: "The time that the Dedicated Host was allocated.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "allows_multiple_instance_types",
				Description: "Indicates whether the Dedicated Host supports multiple instance types of the same instance family",
				Type:        schema.TypeString,
			},
			{
				Name:        "auto_placement",
				Description: "Whether auto-placement is on or off.",
				Type:        schema.TypeString,
			},
			{
				Name:        "availability_zone",
				Description: "The Availability Zone of the Dedicated Host.",
				Type:        schema.TypeString,
			},
			{
				Name:        "availability_zone_id",
				Description: "The ID of the Availability Zone in which the Dedicated Host is allocated.",
				Type:        schema.TypeString,
			},
			{
				Name:        "available_vcpus",
				Description: "The number of vCPUs available for launching instances onto the Dedicated Host.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("AvailableCapacity.AvailableVCpus"),
			},
			{
				Name:        "client_token",
				Description: "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the Dedicated Host.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HostId"),
			},
			{
				Name:        "cores",
				Description: "The number of cores on the Dedicated Host.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HostProperties.Cores"),
			},
			{
				Name:        "instance_family",
				Description: "The instance family supported by the Dedicated Host",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HostProperties.InstanceFamily"),
			},
			{
				Name:        "instance_type",
				Description: "The instance type supported by the Dedicated Host",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HostProperties.InstanceType"),
			},
			{
				Name:        "sockets",
				Description: "The number of sockets on the Dedicated Host.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HostProperties.Sockets"),
			},
			{
				Name:        "total_vcpus",
				Description: "The total number of vCPUs on the Dedicated Host.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("HostProperties.TotalVCpus"),
			},
			{
				Name:        "host_recovery",
				Description: "Indicates whether host recovery is enabled or disabled for the Dedicated Host.",
				Type:        schema.TypeString,
			},
			{
				Name:        "reservation_id",
				Description: "The reservation ID of the Dedicated Host",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("HostReservationId"),
			},
			{
				Name:        "member_of_service_linked_resource_group",
				Description: "Indicates whether the Dedicated Host is in a host resource group",
				Type:        schema.TypeBool,
			},
			{
				Name:        "owner_id",
				Description: "The ID of the Amazon Web Services account that owns the Dedicated Host.",
				Type:        schema.TypeString,
			},
			{
				Name:        "release_time",
				Description: "The time that the Dedicated Host was released.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "state",
				Description: "The Dedicated Host's state.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the Dedicated Host.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_host_available_instance_capacity",
				Description: "Information about the number of instances that can be launched onto the Dedicated Host.",
				Resolver:    fetchEc2HostAvailableInstanceCapacities,
				Columns: []schema.Column{
					{
						Name:        "host_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_hosts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "available_capacity",
						Description: "The number of instances that can be launched onto the Dedicated Host based on the host's available capacity.",
						Type:        schema.TypeInt,
					},
					{
						Name:        "instance_type",
						Description: "The instance type supported by the Dedicated Host.",
						Type:        schema.TypeString,
					},
					{
						Name:        "total_capacity",
						Description: "The total number of instances that can be launched onto the Dedicated Host if there are no instances running on it.",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				Name:        "aws_ec2_host_instances",
				Description: "Describes an instance running on a Dedicated Host.",
				Resolver:    fetchEc2HostInstances,
				Columns: []schema.Column{
					{
						Name:        "host_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_hosts table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "instance_id",
						Description: "The ID of instance that is running on the Dedicated Host.",
						Type:        schema.TypeString,
					},
					{
						Name:        "instance_type",
						Description: "The instance type (for example, m3.medium) of the running instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "owner_id",
						Description: "The ID of the Amazon Web Services account that owns the instance.",
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

func fetchEc2Hosts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2
	input := ec2.DescribeHostsInput{}
	for {
		output, err := svc.DescribeHosts(ctx, &input, func(o *ec2.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Hosts
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
func fetchEc2HostAvailableInstanceCapacities(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	host := parent.Item.(types.Host)
	res <- host.AvailableCapacity.AvailableInstanceCapacity
	return nil
}
func fetchEc2HostInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	host := parent.Item.(types.Host)
	res <- host.Instances
	return nil
}
