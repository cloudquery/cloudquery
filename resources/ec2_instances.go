package resources

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2Instances() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_instances",
		Resolver:     fetchEc2Instances,
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
				Name: "ami_launch_index",
				Type: schema.TypeInt,
			},
			{
				Name: "architecture",
				Type: schema.TypeString,
			},
			{
				Name: "capacity_reservation_id",
				Type: schema.TypeString,
			},
			{
				Name:     "cap_reservation_preference",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CapacityReservationSpecification.CapacityReservationPreference"),
			},
			{
				Name:     "cap_reservation_target_capacity_reservation_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationId"),
			},
			{
				Name:     "cap_reservation_target_capacity_reservation_rg_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationResourceGroupArn"),
			},
			{
				Name: "client_token",
				Type: schema.TypeString,
			},
			{
				Name:     "cpu_options_core_count",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CpuOptions.CoreCount"),
			},
			{
				Name:     "cpu_options_threads_per_core",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("CpuOptions.ThreadsPerCore"),
			},
			{
				Name: "ebs_optimized",
				Type: schema.TypeBool,
			},
			{
				Name: "ena_support",
				Type: schema.TypeBool,
			},
			{
				Name:     "enclave_options_enabled",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EnclaveOptions.Enabled"),
			},
			{
				Name:     "hibernation_options_configured",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("HibernationOptions.Configured"),
			},
			{
				Name: "hypervisor",
				Type: schema.TypeString,
			},
			{
				Name:     "iam_instance_profile_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamInstanceProfile.Arn"),
			},
			{
				Name:     "iam_instance_profile_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamInstanceProfile.Id"),
			},
			{
				Name: "image_id",
				Type: schema.TypeString,
			},
			{
				Name: "instance_id",
				Type: schema.TypeString,
			},
			{
				Name: "instance_lifecycle",
				Type: schema.TypeString,
			},
			{
				Name: "instance_type",
				Type: schema.TypeString,
			},
			{
				Name: "kernel_id",
				Type: schema.TypeString,
			},
			{
				Name: "key_name",
				Type: schema.TypeString,
			},
			{
				Name: "launch_time",
				Type: schema.TypeTimestamp,
			},
			{
				Name:     "metadata_options_http_endpoint",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetadataOptions.HttpEndpoint"),
			},
			{
				Name:     "metadata_options_http_put_response_hop_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MetadataOptions.HttpPutResponseHopLimit"),
			},
			{
				Name:     "metadata_options_http_tokens",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetadataOptions.HttpTokens"),
			},
			{
				Name:     "metadata_options_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("MetadataOptions.State"),
			},
			{
				Name:     "monitoring_state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Monitoring.State"),
			},
			{
				Name: "outpost_arn",
				Type: schema.TypeString,
			},
			{
				Name:     "placement_affinity",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Placement.Affinity"),
			},
			{
				Name:     "placement_availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Placement.AvailabilityZone"),
			},
			{
				Name:     "placement_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Placement.GroupName"),
			},
			{
				Name:     "placement_host_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Placement.HostId"),
			},
			{
				Name:     "placement_host_resource_group_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Placement.HostResourceGroupArn"),
			},
			{
				Name:     "placement_partition_number",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Placement.PartitionNumber"),
			},
			{
				Name:     "placement_spread_domain",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Placement.SpreadDomain"),
			},
			{
				Name:     "placement_tenancy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Placement.Tenancy"),
			},
			{
				Name: "platform",
				Type: schema.TypeString,
			},
			{
				Name: "private_dns_name",
				Type: schema.TypeString,
			},
			{
				Name: "private_ip_address",
				Type: schema.TypeString,
			},
			{
				Name: "public_dns_name",
				Type: schema.TypeString,
			},
			{
				Name: "public_ip_address",
				Type: schema.TypeString,
			},
			{
				Name: "ramdisk_id",
				Type: schema.TypeString,
			},
			{
				Name: "root_device_name",
				Type: schema.TypeString,
			},
			{
				Name: "root_device_type",
				Type: schema.TypeString,
			},
			{
				Name: "source_dest_check",
				Type: schema.TypeBool,
			},
			{
				Name: "spot_instance_request_id",
				Type: schema.TypeString,
			},
			{
				Name: "sriov_net_support",
				Type: schema.TypeString,
			},
			{
				Name:     "state_code",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("State.Code"),
			},
			{
				Name:     "state_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State.Name"),
			},
			{
				Name:     "state_reason_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateReason.Code"),
			},
			{
				Name:     "state_reason_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateReason.Message"),
			},
			{
				Name: "state_transition_reason",
				Type: schema.TypeString,
			},
			{
				Name: "subnet_id",
				Type: schema.TypeString,
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveEc2instanceTags,
			},
			{
				Name: "virtualization_type",
				Type: schema.TypeString,
			},
			{
				Name: "vpc_id",
				Type: schema.TypeString,
			},
		},
		Relations: []*schema.Table{
			{
				Name:     "aws_ec2_instance_block_device_mappings",
				Resolver: fetchEc2InstanceBlockDeviceMappings,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "device_name",
						Type: schema.TypeString,
					},
					{
						Name:     "ebs_attach_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("Ebs.AttachTime"),
					},
					{
						Name:     "ebs_delete_on_termination",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Ebs.DeleteOnTermination"),
					},
					{
						Name:     "ebs_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ebs.Status"),
					},
					{
						Name:     "ebs_volume_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Ebs.VolumeId"),
					},
				},
			},
			{
				Name:     "aws_ec2_instance_elastic_gpu_associations",
				Resolver: fetchEc2InstanceElasticGpuAssociations,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "elastic_gpu_association_id",
						Type: schema.TypeString,
					},
					{
						Name: "elastic_gpu_association_state",
						Type: schema.TypeString,
					},
					{
						Name: "elastic_gpu_association_time",
						Type: schema.TypeString,
					},
					{
						Name: "elastic_gpu_id",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_ec2_instance_elastic_inference_accelerator_associations",
				Resolver: fetchEc2InstanceElasticInferenceAcceleratorAssociations,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "elastic_inference_accelerator_arn",
						Type: schema.TypeString,
					},
					{
						Name: "elastic_inference_accelerator_association_id",
						Type: schema.TypeString,
					},
					{
						Name: "elastic_inference_accelerator_association_state",
						Type: schema.TypeString,
					},
					{
						Name: "elastic_inference_accelerator_association_time",
						Type: schema.TypeTimestamp,
					},
				},
			},
			{
				Name:     "aws_ec2_instance_licenses",
				Resolver: fetchEc2InstanceLicenses,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "license_configuration_arn",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_ec2_instance_network_interfaces",
				Resolver: fetchEc2InstanceNetworkInterfaces,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name:     "association_carrier_ip",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Association.CarrierIp"),
					},
					{
						Name:     "association_ip_owner_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Association.IpOwnerId"),
					},
					{
						Name:     "association_public_dns_name",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Association.PublicDnsName"),
					},
					{
						Name:     "association_public_ip",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Association.PublicIp"),
					},
					{
						Name:     "attachment_attach_time",
						Type:     schema.TypeTimestamp,
						Resolver: schema.PathResolver("Attachment.AttachTime"),
					},
					{
						Name:     "attachment_id",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Attachment.AttachmentId"),
					},
					{
						Name:     "attachment_delete_on_termination",
						Type:     schema.TypeBool,
						Resolver: schema.PathResolver("Attachment.DeleteOnTermination"),
					},
					{
						Name:     "attachment_device_index",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Attachment.DeviceIndex"),
					},
					{
						Name:     "attachment_network_card_index",
						Type:     schema.TypeInt,
						Resolver: schema.PathResolver("Attachment.NetworkCardIndex"),
					},
					{
						Name:     "attachment_status",
						Type:     schema.TypeString,
						Resolver: schema.PathResolver("Attachment.Status"),
					},
					{
						Name: "description",
						Type: schema.TypeString,
					},
					{
						Name: "interface_type",
						Type: schema.TypeString,
					},
					{
						Name: "mac_address",
						Type: schema.TypeString,
					},
					{
						Name: "network_interface_id",
						Type: schema.TypeString,
					},
					{
						Name: "owner_id",
						Type: schema.TypeString,
					},
					{
						Name: "private_dns_name",
						Type: schema.TypeString,
					},
					{
						Name: "private_ip_address",
						Type: schema.TypeString,
					},
					{
						Name: "source_dest_check",
						Type: schema.TypeBool,
					},
					{
						Name: "status",
						Type: schema.TypeString,
					},
					{
						Name: "subnet_id",
						Type: schema.TypeString,
					},
					{
						Name: "vpc_id",
						Type: schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:     "aws_ec2_instance_network_interface_groups",
						Resolver: fetchEc2InstanceNetworkInterfaceGroups,
						Columns: []schema.Column{
							{
								Name:     "instance_network_interface_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "group_id",
								Type: schema.TypeString,
							},
							{
								Name: "group_name",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_ec2_instance_network_interface_ipv6_addresses",
						Resolver: fetchEc2InstanceNetworkInterfaceIpv6Addresses,
						Columns: []schema.Column{
							{
								Name:     "instance_network_interface_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name: "ipv6_address",
								Type: schema.TypeString,
							},
						},
					},
					{
						Name:     "aws_ec2_instance_network_interface_instance_private_ip_addresses",
						Resolver: fetchEc2InstanceNetworkInterfaceInstancePrivateIpAddresses,
						Columns: []schema.Column{
							{
								Name:     "instance_network_interface_id",
								Type:     schema.TypeUUID,
								Resolver: schema.ParentIdResolver,
							},
							{
								Name:     "association_carrier_ip",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Association.CarrierIp"),
							},
							{
								Name:     "association_ip_owner_id",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Association.IpOwnerId"),
							},
							{
								Name:     "association_public_dns_name",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Association.PublicDnsName"),
							},
							{
								Name:     "association_public_ip",
								Type:     schema.TypeString,
								Resolver: schema.PathResolver("Association.PublicIp"),
							},
							{
								Name:     "is_primary",
								Type:     schema.TypeBool,
								Resolver: schema.PathResolver("Primary"),
							},
							{
								Name: "private_dns_name",
								Type: schema.TypeString,
							},
							{
								Name: "private_ip_address",
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:     "aws_ec2_instance_product_codes",
				Resolver: fetchEc2InstanceProductCodes,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "product_code_id",
						Type: schema.TypeString,
					},
					{
						Name: "product_code_type",
						Type: schema.TypeString,
					},
				},
			},
			{
				Name:     "aws_ec2_instance_security_groups",
				Resolver: fetchEc2InstanceSecurityGroups,
				Columns: []schema.Column{
					{
						Name:     "instance_id",
						Type:     schema.TypeUUID,
						Resolver: schema.ParentIdResolver,
					},
					{
						Name: "group_id",
						Type: schema.TypeString,
					},
					{
						Name: "group_name",
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================
func fetchEc2Instances(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2

	response, err := svc.DescribeInstances(ctx, &ec2.DescribeInstancesInput{}, func(o *ec2.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return err
	}

	for _, reservation := range response.Reservations {
		res <- reservation.Instances
	}

	return nil
}
func resolveEc2instanceTags(_ context.Context, _ schema.ClientMeta, resource *schema.Resource, _ schema.Column) error {
	r := resource.Item.(types.Instance)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	resource.Set("tags", tags)
	return nil
}
func fetchEc2InstanceBlockDeviceMappings(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.BlockDeviceMappings
	return nil
}
func fetchEc2InstanceElasticGpuAssociations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.ElasticGpuAssociations
	return nil
}
func fetchEc2InstanceElasticInferenceAcceleratorAssociations(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.ElasticInferenceAcceleratorAssociations
	return nil
}
func fetchEc2InstanceLicenses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.Licenses
	return nil
}
func fetchEc2InstanceNetworkInterfaces(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.NetworkInterfaces
	return nil
}
func fetchEc2InstanceNetworkInterfaceGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instanceNetworkInterface, ok := parent.Item.(types.InstanceNetworkInterface)
	if !ok {
		return fmt.Errorf("not ec2 instance network interface")
	}
	res <- instanceNetworkInterface.Groups
	return nil
}
func fetchEc2InstanceNetworkInterfaceIpv6Addresses(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instanceNetworkInterface, ok := parent.Item.(types.InstanceNetworkInterface)
	if !ok {
		return fmt.Errorf("not ec2 instance network interface")
	}
	res <- instanceNetworkInterface.Ipv6Addresses
	return nil
}
func fetchEc2InstanceNetworkInterfaceInstancePrivateIpAddresses(__ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instanceNetworkInterface, ok := parent.Item.(types.InstanceNetworkInterface)
	if !ok {
		return fmt.Errorf("not ec2 instance network interface")
	}
	res <- instanceNetworkInterface.PrivateIpAddresses
	return nil
}
func fetchEc2InstanceProductCodes(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.ProductCodes
	return nil
}
func fetchEc2InstanceSecurityGroups(_ context.Context, _ schema.ClientMeta, parent *schema.Resource, res chan interface{}) error {
	instance, ok := parent.Item.(types.Instance)
	if !ok {
		return fmt.Errorf("not ec2 instance")
	}
	res <- instance.SecurityGroups
	return nil
}
