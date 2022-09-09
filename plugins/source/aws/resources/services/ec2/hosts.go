package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Hosts() *schema.Table {
	return &schema.Table{
		Name:          "aws_ec2_hosts",
		Description:   "Describes the properties of the Dedicated Host.",
		Resolver:      fetchEc2Hosts,
		Multiplex:     client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreInTests: true,
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
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
			{
				Name:     "available_capacity",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AvailableCapacity"),
			},
			{
				Name:        "instances",
				Type:        schema.TypeJSON,
				Description: "Describes an instance running on a Dedicated Host.",
				Resolver:    schema.PathResolver("Instances"),
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
			return err
		}
		res <- output.Hosts
		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}
	return nil
}
