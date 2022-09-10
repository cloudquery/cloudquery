package ec2

import (
	"context"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Ec2Instances() *schema.Table {
	return &schema.Table{
		Name:        "aws_ec2_instances",
		Description: "Describes an instance.",
		Resolver:    fetchEc2Instances,
		Multiplex:   client.ServiceAccountRegionMultiplexer("ec2"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
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
				Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
					return []string{"instance", *resource.Item.(types.Instance).InstanceId}, nil
				}),
			},
			{
				Name:          "state_transition_reason_time",
				Type:          schema.TypeTimestamp,
				Resolver:      resolveEc2InstanceStateTransitionReasonTime,
				IgnoreInTests: true,
			},
			{
				Name:          "ami_launch_index",
				Description:   "The AMI launch index, which can be used to find this instance in the launch group.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:        "architecture",
				Description: "The architecture of the image.",
				Type:        schema.TypeString,
			},
			{
				Name:        "boot_mode",
				Description: "The boot mode of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:          "capacity_reservation_id",
				Description:   "The ID of the Capacity Reservation.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "capacity_reservation_specification",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("CapacityReservationSpecification"),
				IgnoreInTests: true,
			},
			{
				Name:        "client_token",
				Description: "The idempotency token you provided when you launched the instance, if applicable.",
				Type:        schema.TypeString,
			},
			{
				Name:     "cpu_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("CpuOptions"),
			},
			{
				Name:        "ebs_optimized",
				Description: "Indicates whether the instance is optimized for Amazon EBS I/O",
				Type:        schema.TypeBool,
			},
			{
				Name:        "ena_support",
				Description: "Specifies whether enhanced networking with ENA is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:        "enclave_options",
				Type: 			schema.TypeJSON,
			},
			{
				Name:        "hibernation_options",
				Type:        schema.TypeJSON,
			},
			{
				Name:        "hypervisor",
				Description: "The hypervisor type of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:          "iam_instance_profile",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("IamInstanceProfile"),
			},
			{
				Name:        "image_id",
				Description: "The ID of the AMI used to launch the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:            "id",
				Description:     "The ID of the instance.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("InstanceId"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "instance_lifecycle",
				Description: "Indicates whether this is a Spot Instance or a Scheduled Instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "instance_type",
				Description: "The instance type.",
				Type:        schema.TypeString,
			},
			{
				Name:          "kernel_id",
				Description:   "The kernel associated with this instance, if applicable.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "key_name",
				Description:   "The name of the key pair, if this instance was launched with an associated key pair.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "launch_time",
				Description: "The time the instance was launched.",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "licenses",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Licenses"),
			},
			{
				Name:     "metadata_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetadataOptions"),
			},
			{
				Name:     "monitoring",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Monitoring"),
			},
			{
				Name:          "outpost_arn",
				Description:   "The Amazon Resource Name (ARN) of the Outpost.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "placement",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("Placement"),
				IgnoreInTests: true,
			},
			{
				Name:        "platform",
				Description: "The value is Windows for Windows instances; otherwise blank.",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_dns_name",
				Description: "(IPv4 only) The private DNS hostname name assigned to the instance",
				Type:        schema.TypeString,
			},
			{
				Name:        "private_ip_address",
				Description: "The private IPv4 address assigned to the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "public_dns_name",
				Description: "(IPv4 only) The public DNS name assigned to the instance",
				Type:        schema.TypeString,
			},
			{
				Name:          "public_ip_address",
				Description:   "The public IPv4 address, or the Carrier IP address assigned to the instance, if applicable",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "ramdisk_id",
				Description:   "The RAM disk associated with this instance, if applicable.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "root_device_name",
				Description: "The device name of the root device volume (for example, /dev/sda1).",
				Type:        schema.TypeString,
			},
			{
				Name:        "root_device_type",
				Description: "The root device type used by the AMI",
				Type:        schema.TypeString,
			},
			{
				Name:        "source_dest_check",
				Description: "Indicates whether source/destination checking is enabled.",
				Type:        schema.TypeBool,
			},
			{
				Name:          "spot_instance_request_id",
				Description:   "If the request is a Spot Instance request, the ID of the request.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "sriov_net_support",
				Description:   "Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:     "state",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:          "state_reason",
				Type:          schema.TypeJSON,
				Resolver:      schema.PathResolver("StateReason"),
			},
			{
				Name:        "state_transition_reason",
				Description: "The reason for the most recent state transition",
				Type:        schema.TypeString,
			},
			{
				Name:        "subnet_id",
				Description: "[EC2-VPC] The ID of the subnet in which the instance is running.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Any tags assigned to the instance.",
				Type:        schema.TypeJSON,
				Resolver:    client.ResolveTags,
			},
			{
				Name:        "virtualization_type",
				Description: "The virtualization type of the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "vpc_id",
				Description: "The ID of the VPC in which the instance is running.",
				Type:        schema.TypeString,
			},
			{
				Name:        "block_device_mappings",
				Description: "Describes a block device mapping.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("BlockDeviceMappings"),
			},
			{
				Name:        "elastic_gpu_associations",
				Description: "Describes the association between an instance and an Elastic Graphics accelerator.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ElasticGpuAssociations"),
			},
			{
				Name:        "elastic_inference_accelerator_associations",
				Description: "Describes the association between an instance and an elastic inference accelerator.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ElasticInferenceAcceleratorAssociations"),
			},
			{
				Name:        "network_interfaces",
				Description: "Describes a network interface.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("NetworkInterfaces"),
			},
			{
				Name:        "product_codes",
				Description: "Describes a product code.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("ProductCodes"),
			},
			{
				Name:        "security_groups",
				Description: "Describes a security group.",
				Type:        schema.TypeJSON,
				Resolver:    schema.PathResolver("SecurityGroups"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchEc2Instances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	var config ec2.DescribeInstancesInput
	c := meta.(*client.Client)
	svc := c.Services().EC2
	for {
		output, err := svc.DescribeInstances(ctx, &config, func(options *ec2.Options) {
			options.Region = c.Region
		})
		if err != nil {
			return err
		}
		for _, reservation := range output.Reservations {
			res <- reservation.Instances
		}

		if aws.ToString(output.NextToken) == "" {
			break
		}
		config.NextToken = output.NextToken
	}
	return nil
}
func resolveEc2InstanceStateTransitionReasonTime(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(types.Instance)
	if instance.StateTransitionReason == nil {
		return nil
	}
	re := regexp.MustCompile(`\((.*)\)`)
	match := re.FindStringSubmatch(*instance.StateTransitionReason)
	if len(match) < 2 {
		// failed to get time from message
		return nil
	}
	const layout = "2006-01-02 15:04:05 MST"
	tm, err := time.Parse(layout, match[1])
	if err != nil {
		// failed to parse last transition time
		return nil
	}
	return resource.Set(c.Name, tm)
}
