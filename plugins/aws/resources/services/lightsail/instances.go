package lightsail

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

//go:generate cq-gen --resource instances --config gen.hcl --output .
func Instances() *schema.Table {
	return &schema.Table{
		Name:         "aws_lightsail_instances",
		Description:  "Describes an instance (a virtual private server)",
		Resolver:     fetchLightsailInstances,
		Multiplex:    client.ServiceAccountRegionMultiplexer("lightsail"),
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
				Name:     "access_details",
				Type:     schema.TypeJSON,
				Resolver: ResolveLightsailInstanceAccessDetails,
			},
			{
				Name:        "arn",
				Description: "The Amazon Resource Name (ARN) of the instance (eg, arn:aws:lightsail:us-east-2:123456789101:Instance/244ad76f-8aad-4741-809f-12345EXAMPLE)",
				Type:        schema.TypeString,
			},
			{
				Name:        "blueprint_id",
				Description: "The blueprint ID (eg, os_amlinux_2016_03)",
				Type:        schema.TypeString,
			},
			{
				Name:        "blueprint_name",
				Description: "The friendly name of the blueprint (eg, Amazon Linux)",
				Type:        schema.TypeString,
			},
			{
				Name:        "bundle_id",
				Description: "The bundle for the instance (eg, micro_1_0)",
				Type:        schema.TypeString,
			},
			{
				Name:        "created_at",
				Description: "The timestamp when the instance was created (eg, 147973490917) in Unix time format",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "hardware_cpu_count",
				Description: "The number of vCPUs the instance has",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Hardware.CpuCount"),
			},
			{
				Name:        "hardware_ram_size_in_gb",
				Description: "The amount of RAM in GB on the instance (eg, 10)",
				Type:        schema.TypeFloat,
				Resolver:    schema.PathResolver("Hardware.RamSizeInGb"),
			},
			{
				Name:        "ip_address_type",
				Description: "The IP address type of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "ipv6_addresses",
				Description: "The IPv6 addresses of the instance",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "is_static_ip",
				Description: "A Boolean value indicating whether this instance has a static IP assigned to it",
				Type:        schema.TypeBool,
			},
			{
				Name:        "location_availability_zone",
				Description: "The Availability Zone",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.AvailabilityZone"),
			},
			{
				Name:        "location_region_name",
				Description: "The AWS Region name",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Location.RegionName"),
			},
			{
				Name:        "name",
				Description: "The name the user gave the instance (eg, Amazon_Linux-1GB-Ohio-1)",
				Type:        schema.TypeString,
			},
			{
				Name:        "networking_monthly_transfer_gb_per_month_allocated",
				Description: "The amount allocated per month (in GB)",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("Networking.MonthlyTransfer.GbPerMonthAllocated"),
			},
			{
				Name:        "private_ip_address",
				Description: "The private IP address of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_ip_address",
				Description: "The public IP address of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "resource_type",
				Description: "The type of resource (usually Instance)",
				Type:        schema.TypeString,
			},
			{
				Name:        "ssh_key_name",
				Description: "The name of the SSH key being used to connect to the instance (eg, LightsailDefaultKeyPair)",
				Type:        schema.TypeString,
			},
			{
				Name:        "state_code",
				Description: "The status code for the instance",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("State.Code"),
			},
			{
				Name:        "state_name",
				Description: "The state of the instance (eg, running or pending)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State.Name"),
			},
			{
				Name:        "support_code",
				Description: "The support code",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "The tag keys and optional values for the resource",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "username",
				Description: "The user name for connecting to the instance (eg, ec2-user)",
				Type:        schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:          "aws_lightsail_instance_add_ons",
				Description:   "Describes an add-on that is enabled for an Amazon Lightsail resource",
				Resolver:      fetchLightsailInstanceAddOns,
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "name",
						Description: "The name of the add-on",
						Type:        schema.TypeString,
					},
					{
						Name:        "next_snapshot_time_of_day",
						Description: "The next daily time an automatic snapshot will be created",
						Type:        schema.TypeString,
					},
					{
						Name:        "snapshot_time_of_day",
						Description: "The daily time when an automatic snapshot is created",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the add-on",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_lightsail_instance_hardware_disks",
				Description: "Describes a block storage disk",
				Resolver:    fetchLightsailInstanceHardwareDisks,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) of the disk",
						Type:        schema.TypeString,
					},
					{
						Name:        "attached_to",
						Description: "The resources to which the disk is attached",
						Type:        schema.TypeString,
					},
					{
						Name:        "created_at",
						Description: "The date when the disk was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:          "gb_in_use",
						Description:   "(Deprecated) The number of GB in use by the disk",
						Type:          schema.TypeInt,
						IgnoreInTests: true,
					},
					{
						Name:        "iops",
						Description: "The input/output operations per second (IOPS) of the disk",
						Type:        schema.TypeInt,
					},
					{
						Name:        "is_attached",
						Description: "A Boolean value indicating whether the disk is attached",
						Type:        schema.TypeBool,
					},
					{
						Name:        "is_system_disk",
						Description: "A Boolean value indicating whether this disk is a system disk (has an operating system loaded on it)",
						Type:        schema.TypeBool,
					},
					{
						Name:        "location_availability_zone",
						Description: "The Availability Zone",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.AvailabilityZone"),
					},
					{
						Name:        "location_region_name",
						Description: "The AWS Region name",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Location.RegionName"),
					},
					{
						Name:        "name",
						Description: "The unique name of the disk",
						Type:        schema.TypeString,
					},
					{
						Name:        "path",
						Description: "The disk path",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_type",
						Description: "The Lightsail resource type (eg, Disk)",
						Type:        schema.TypeString,
					},
					{
						Name:        "size_in_gb",
						Description: "The size of the disk in GB",
						Type:        schema.TypeInt,
					},
					{
						Name:        "state",
						Description: "Describes the status of the disk",
						Type:        schema.TypeString,
					},
					{
						Name:        "support_code",
						Description: "The support code",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "The tag keys and optional values for the resource",
						Type:        schema.TypeJSON,
						Resolver:    client.ResolveTags,
					},
				},
				Relations: []*schema.Table{
					{
						Name:          "aws_lightsail_instance_hardware_disk_add_ons",
						Description:   "Describes an add-on that is enabled for an Amazon Lightsail resource",
						Resolver:      fetchLightsailInstanceHardwareDiskAddOns,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "instance_hardware_disk_cq_id",
								Description: "Unique CloudQuery ID of aws_lightsail_instance_hardware_disks table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "name",
								Description: "The name of the add-on",
								Type:        schema.TypeString,
							},
							{
								Name:        "next_snapshot_time_of_day",
								Description: "The next daily time an automatic snapshot will be created",
								Type:        schema.TypeString,
							},
							{
								Name:        "snapshot_time_of_day",
								Description: "The daily time when an automatic snapshot is created",
								Type:        schema.TypeString,
							},
							{
								Name:        "status",
								Description: "The status of the add-on",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:        "aws_lightsail_instance_networking_ports",
				Description: "Describes information about ports for an Amazon Lightsail instance",
				Resolver:    fetchLightsailInstanceNetworkingPorts,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "access_direction",
						Description: "The access direction (inbound or outbound)",
						Type:        schema.TypeString,
					},
					{
						Name:        "access_from",
						Description: "The location from which access is allowed",
						Type:        schema.TypeString,
					},
					{
						Name:        "access_type",
						Description: "The type of access (Public or Private)",
						Type:        schema.TypeString,
					},
					{
						Name:        "cidr_list_aliases",
						Description: "An alias that defines access for a preconfigured range of IP addresses",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "cidrs",
						Description: "The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "common_name",
						Description: "The common name of the port information",
						Type:        schema.TypeString,
					},
					{
						Name:        "from_port",
						Description: "The first port in a range of open ports on an instance",
						Type:        schema.TypeInt,
					},
					{
						Name:        "ipv6_cidrs",
						Description: "The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "protocol",
						Description: "The IP protocol name",
						Type:        schema.TypeString,
					},
					{
						Name:        "to_port",
						Description: "The last port in a range of open ports on an instance",
						Type:        schema.TypeInt,
					},
				},
			},
			{
				Name:        "aws_lightsail_instance_port_states",
				Description: "Describes open ports on an instance, the IP addresses allowed to connect to the instance through the ports, and the protocol",
				Resolver:    fetchLightsailInstancePortStates,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_lightsail_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "cidr_list_aliases",
						Description: "An alias that defines access for a preconfigured range of IP addresses",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "cidrs",
						Description: "The IPv4 address, or range of IPv4 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "from_port",
						Description: "The first port in a range of open ports on an instance",
						Type:        schema.TypeInt,
					},
					{
						Name:        "ipv6_cidrs",
						Description: "The IPv6 address, or range of IPv6 addresses (in CIDR notation) that are allowed to connect to an instance through the ports, and the protocol",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "protocol",
						Description: "The IP protocol name",
						Type:        schema.TypeString,
					},
					{
						Name:        "state",
						Description: "Specifies whether the instance port is open or closed",
						Type:        schema.TypeString,
					},
					{
						Name:        "to_port",
						Description: "The last port in a range of open ports on an instance",
						Type:        schema.TypeInt,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchLightsailInstances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().Lightsail
	input := lightsail.GetInstancesInput{}
	for {
		output, err := svc.GetInstances(ctx, &input, func(o *lightsail.Options) {
			o.Region = c.Region
		})
		if err != nil {
			return diag.WrapError(err)
		}
		res <- output.Instances

		if aws.ToString(output.NextPageToken) == "" {
			break
		}
		input.PageToken = output.NextPageToken
	}
	return nil
}
func ResolveLightsailInstanceAccessDetails(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services().Lightsail
	input := lightsail.GetInstanceAccessDetailsInput{InstanceName: r.Name}
	output, err := svc.GetInstanceAccessDetails(ctx, &input, func(o *lightsail.Options) {
		o.Region = cli.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}
	j, err := json.Marshal(output.AccessDetails)
	if err != nil {
		return diag.WrapError(err)
	}
	return diag.WrapError(resource.Set(c.Name, j))
}
func fetchLightsailInstanceAddOns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	res <- instance.AddOns
	return nil
}
func fetchLightsailInstanceHardwareDisks(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	if instance.Hardware == nil {
		return nil
	}
	res <- instance.Hardware.Disks
	return nil
}
func fetchLightsailInstanceHardwareDiskAddOns(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	disk := parent.Item.(types.Disk)
	res <- disk.AddOns
	return nil
}
func fetchLightsailInstanceNetworkingPorts(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	if instance.Networking == nil {
		return nil
	}
	res <- instance.Networking.Ports
	return nil
}
func fetchLightsailInstancePortStates(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	r := parent.Item.(types.Instance)
	cli := meta.(*client.Client)
	svc := cli.Services().Lightsail
	input := lightsail.GetInstancePortStatesInput{InstanceName: r.Name}
	output, err := svc.GetInstancePortStates(ctx, &input, func(o *lightsail.Options) {
		o.Region = cli.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}

	res <- output.PortStates
	return nil
}
