package ec2

import (
	"context"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/cloudquery/cq-provider-aws/client"

	"github.com/cloudquery/cq-provider-sdk/provider/diag"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
)

func Ec2Instances() *schema.Table {
	return &schema.Table{
		Name:         "aws_ec2_instances",
		Description:  "Describes an instance.",
		Resolver:     fetchEc2Instances,
		Multiplex:    client.ServiceAccountRegionMultiplexer("ec2"),
		IgnoreError:  client.IgnoreAccessDeniedServiceDisabled,
		DeleteFilter: client.DeleteAccountRegionFilter,
		Options:      schema.TableCreationOptions{PrimaryKeys: []string{"account_id", "id"}},
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
				Name:        "cap_reservation_preference",
				Description: "Describes the instance's Capacity Reservation preferences",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("CapacityReservationSpecification.CapacityReservationPreference"),
			},
			{
				Name:          "cap_reservation_target_capacity_reservation_id",
				Description:   "The ID of the targeted Capacity Reservation.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationId"),
				IgnoreInTests: true,
			},
			{
				Name:          "cap_reservation_target_capacity_reservation_rg_arn",
				Description:   "The ARN of the targeted Capacity Reservation group.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("CapacityReservationSpecification.CapacityReservationTarget.CapacityReservationResourceGroupArn"),
				IgnoreInTests: true,
			},
			{
				Name:        "client_token",
				Description: "The idempotency token you provided when you launched the instance, if applicable.",
				Type:        schema.TypeString,
			},
			{
				Name:        "cpu_options_core_count",
				Description: "The number of CPU cores for the instance.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CpuOptions.CoreCount"),
			},
			{
				Name:        "cpu_options_threads_per_core",
				Description: "The number of threads per CPU core.",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("CpuOptions.ThreadsPerCore"),
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
				Name:        "enclave_options_enabled",
				Description: "If this parameter is set to true, the instance is enabled for Amazon Web Services Nitro Enclaves; otherwise, it is not enabled for Amazon Web Services Nitro Enclaves.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("EnclaveOptions.Enabled"),
			},
			{
				Name:        "hibernation_options_configured",
				Description: "If this parameter is set to true, your instance is enabled for hibernation; otherwise, it is not enabled for hibernation.",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("HibernationOptions.Configured"),
			},
			{
				Name:        "hypervisor",
				Description: "The hypervisor type of the instance",
				Type:        schema.TypeString,
			},
			{
				Name:          "iam_instance_profile_arn",
				Description:   "The Amazon Resource Name (ARN) of the instance profile.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("IamInstanceProfile.Arn"),
				IgnoreInTests: true,
			},
			{
				Name:          "iam_instance_profile_id",
				Description:   "The ID of the instance profile.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("IamInstanceProfile.Id"),
				IgnoreInTests: true,
			},
			{
				Name:        "image_id",
				Description: "The ID of the AMI used to launch the instance.",
				Type:        schema.TypeString,
			},
			{
				Name:        "id",
				Description: "The ID of the instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("InstanceId"),
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
				Name:        "licenses",
				Description: "The license configurations.",
				Type:        schema.TypeStringArray,
				Resolver:    resolveEc2InstancesLicenses,
			},
			{
				Name:        "metadata_options_http_endpoint",
				Description: "This parameter enables or disables the HTTP metadata endpoint on your instances. If the parameter is not specified, the default state is enabled",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.HttpEndpoint"),
			},
			{
				Name:        "metadata_options_http_protocol_ipv6",
				Description: "Whether or not the IPv6 endpoint for the instance metadata service is enabled or disabled.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.HttpProtocolIpv6"),
			},
			{
				Name:        "metadata_options_http_put_response_hop_limit",
				Description: "The desired HTTP PUT response hop limit for instance metadata requests",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("MetadataOptions.HttpPutResponseHopLimit"),
			},
			{
				Name:        "metadata_options_http_tokens",
				Description: "The state of token usage for your instance metadata requests",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.HttpTokens"),
			},
			{
				Name:        "metadata_options_state",
				Description: "The state of the metadata option changes",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("MetadataOptions.State"),
			},
			{
				Name:        "monitoring_state",
				Description: "Indicates whether detailed monitoring is enabled",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Monitoring.State"),
			},
			{
				Name:          "outpost_arn",
				Description:   "The Amazon Resource Name (ARN) of the Outpost.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:          "placement_affinity",
				Description:   "The affinity setting for the instance on the Dedicated Host",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Placement.Affinity"),
				IgnoreInTests: true,
			},
			{
				Name:        "placement_availability_zone",
				Description: "The Availability Zone of the instance",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.AvailabilityZone"),
			},
			{
				Name:        "placement_group_name",
				Description: "The name of the placement group the instance is in.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.GroupName"),
			},
			{
				Name:          "placement_host_id",
				Description:   "The ID of the Dedicated Host on which the instance resides",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Placement.HostId"),
				IgnoreInTests: true,
			},
			{
				Name:          "placement_host_resource_group_arn",
				Description:   "The ARN of the host resource group in which to launch the instances",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Placement.HostResourceGroupArn"),
				IgnoreInTests: true,
			},
			{
				Name:          "placement_partition_number",
				Description:   "The number of the partition the instance is in",
				Type:          schema.TypeInt,
				Resolver:      schema.PathResolver("Placement.PartitionNumber"),
				IgnoreInTests: true,
			},
			{
				Name:          "placement_spread_domain",
				Description:   "Reserved for future use",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("Placement.SpreadDomain"),
				IgnoreInTests: true,
			},
			{
				Name:        "placement_tenancy",
				Description: "The tenancy of the instance (if the instance is running in a VPC)",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("Placement.Tenancy"),
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
				Name:        "state_code",
				Description: "The state of the instance as a 16-bit unsigned integer",
				Type:        schema.TypeInt,
				Resolver:    schema.PathResolver("State.Code"),
			},
			{
				Name:        "state_name",
				Description: "The current state of the instance.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("State.Name"),
			},
			{
				Name:          "state_reason_code",
				Description:   "The reason code for the state change.",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StateReason.Code"),
				IgnoreInTests: true,
			},
			{
				Name:          "state_reason_message",
				Description:   "The message for the state change.  * Server.InsufficientInstanceCapacity: There was insufficient capacity available to satisfy the launch request.  * Server.InternalError: An internal error caused the instance to terminate during launch.  * Server.ScheduledStop: The instance was stopped due to a scheduled retirement.  * Server.SpotInstanceShutdown: The instance was stopped because the number of Spot requests with a maximum price equal to or higher than the Spot price exceeded available capacity or because of an increase in the Spot price.  * Server.SpotInstanceTermination: The instance was terminated because the number of Spot requests with a maximum price equal to or higher than the Spot price exceeded available capacity or because of an increase in the Spot price.  * Client.InstanceInitiatedShutdown: The instance was shut down using the shutdown -h command from the instance.  * Client.InstanceTerminated: The instance was terminated or rebooted during AMI creation.  * Client.InternalError: A client error caused the instance to terminate during launch.  * Client.InvalidSnapshot.NotFound: The specified snapshot was not found.  * Client.UserInitiatedHibernate: Hibernation was initiated on the instance.  * Client.UserInitiatedShutdown: The instance was shut down using the Amazon EC2 API.  * Client.VolumeLimitExceeded: The limit on the number of EBS volumes or total storage was exceeded",
				Type:          schema.TypeString,
				Resolver:      schema.PathResolver("StateReason.Message"),
				IgnoreInTests: true,
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
				Resolver:    resolveEc2InstancesTags,
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
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_ec2_instance_block_device_mappings",
				Description: "Describes a block device mapping.",
				Resolver:    fetchEc2InstanceBlockDeviceMappings,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "ebs_volume_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "device_name",
						Description: "The device name (for example, /dev/sdh or xvdh).",
						Type:        schema.TypeString,
					},
					{
						Name:        "ebs_attach_time",
						Description: "The time stamp when the attachment initiated.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("Ebs.AttachTime"),
					},
					{
						Name:        "ebs_delete_on_termination",
						Description: "Indicates whether the volume is deleted on instance termination.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Ebs.DeleteOnTermination"),
					},
					{
						Name:        "ebs_status",
						Description: "The attachment state.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.Status"),
					},
					{
						Name:        "ebs_volume_id",
						Description: "The ID of the EBS volume.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Ebs.VolumeId"),
					},
				},
			},
			{
				Name:          "aws_ec2_instance_elastic_gpu_associations",
				Description:   "Describes the association between an instance and an Elastic Graphics accelerator.",
				Resolver:      fetchEc2InstanceElasticGpuAssociations,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "elastic_gpu_association_id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "elastic_gpu_association_id",
						Description: "The ID of the association.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_gpu_association_state",
						Description: "The state of the association between the instance and the Elastic Graphics accelerator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_gpu_association_time",
						Description: "The time the Elastic Graphics accelerator was associated with the instance.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_gpu_id",
						Description: "The ID of the Elastic Graphics accelerator.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:          "aws_ec2_instance_elastic_inference_accelerator_associations",
				Description:   "Describes the association between an instance and an elastic inference accelerator.",
				Resolver:      fetchEc2InstanceElasticInferenceAcceleratorAssociations,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "elastic_inference_accelerator_association_id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "elastic_inference_accelerator_arn",
						Description: "The Amazon Resource Name (ARN) of the elastic inference accelerator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_inference_accelerator_association_id",
						Description: "The ID of the association.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_inference_accelerator_association_state",
						Description: "The state of the elastic inference accelerator.",
						Type:        schema.TypeString,
					},
					{
						Name:        "elastic_inference_accelerator_association_time",
						Description: "The time at which the elastic inference accelerator is associated with an instance.",
						Type:        schema.TypeTimestamp,
					},
				},
			},
			{
				Name:        "aws_ec2_instance_network_interfaces",
				Description: "Describes a network interface.",
				Resolver:    fetchEc2InstanceNetworkInterfaces,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "network_interface_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "arn",
						Description: "The Amazon Resource Name (ARN) for the resource.",
						Type:        schema.TypeString,
						Resolver: client.ResolveARN(client.EC2Service, func(resource *schema.Resource) ([]string, error) {
							return []string{"network-interface", *resource.Item.(types.InstanceNetworkInterface).NetworkInterfaceId}, nil
						}),
					},
					{
						Name:          "association_carrier_ip",
						Description:   "The carrier IP address associated with the network interface.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Association.CarrierIp"),
						IgnoreInTests: true,
					},
					{
						Name:          "association_ip_owner_id",
						Description:   "The ID of the owner of the Elastic IP address.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Association.IpOwnerId"),
						IgnoreInTests: true,
					},
					{
						Name:          "association_public_dns_name",
						Description:   "The public DNS name.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Association.PublicDnsName"),
						IgnoreInTests: true,
					},
					{
						Name:          "association_public_ip",
						Description:   "The public IP address or Elastic IP address bound to the network interface.",
						Type:          schema.TypeString,
						Resolver:      schema.PathResolver("Association.PublicIp"),
						IgnoreInTests: true,
					},
					{
						Name:        "attachment_attach_time",
						Description: "The time stamp when the attachment initiated.",
						Type:        schema.TypeTimestamp,
						Resolver:    schema.PathResolver("Attachment.AttachTime"),
					},
					{
						Name:        "attachment_id",
						Description: "The ID of the network interface attachment.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Attachment.AttachmentId"),
					},
					{
						Name:        "attachment_delete_on_termination",
						Description: "Indicates whether the network interface is deleted when the instance is terminated.",
						Type:        schema.TypeBool,
						Resolver:    schema.PathResolver("Attachment.DeleteOnTermination"),
					},
					{
						Name:        "attachment_device_index",
						Description: "The index of the device on the instance for the network interface attachment.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Attachment.DeviceIndex"),
					},
					{
						Name:        "attachment_network_card_index",
						Description: "The index of the network card.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("Attachment.NetworkCardIndex"),
					},
					{
						Name:        "attachment_status",
						Description: "The attachment state.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("Attachment.Status"),
					},
					{
						Name:        "description",
						Description: "The description.",
						Type:        schema.TypeString,
					},
					{
						Name:        "interface_type",
						Description: "Describes the type of network interface",
						Type:        schema.TypeString,
					},
					{
						Name:        "ipv4_prefixes",
						Description: "The IPv4 delegated prefixes that are assigned to the network interface.",
						Type:        schema.TypeStringArray,
						Resolver:    resolveEc2InstanceNetworkInterfacesIpv4Prefixes,
					},
					{
						Name:        "ipv6_prefixes",
						Description: "The IPv6 delegated prefixes that are assigned to the network interface.",
						Type:        schema.TypeStringArray,
						Resolver:    resolveEc2InstanceNetworkInterfacesIpv6Prefixes,
					},
					{
						Name:        "mac_address",
						Description: "The MAC address.",
						Type:        schema.TypeString,
					},
					{
						Name:        "network_interface_id",
						Description: "The ID of the network interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "owner_id",
						Description: "The ID of the Amazon Web Services account that created the network interface.",
						Type:        schema.TypeString,
					},
					{
						Name:          "private_dns_name",
						Description:   "The private DNS name.",
						Type:          schema.TypeString,
						IgnoreInTests: true,
					},
					{
						Name:        "private_ip_address",
						Description: "The IPv4 address of the network interface within the subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "source_dest_check",
						Description: "Indicates whether source/destination checking is enabled.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "status",
						Description: "The status of the network interface.",
						Type:        schema.TypeString,
					},
					{
						Name:        "subnet_id",
						Description: "The ID of the subnet.",
						Type:        schema.TypeString,
					},
					{
						Name:        "vpc_id",
						Description: "The ID of the VPC.",
						Type:        schema.TypeString,
					},
				},
				Relations: []*schema.Table{
					{
						Name:        "aws_ec2_instance_network_interface_groups",
						Description: "Describes a security group.",
						Resolver:    fetchEc2InstanceNetworkInterfaceGroups,
						Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_network_interface_cq_id", "group_id"}},
						Columns: []schema.Column{
							{
								Name:        "instance_network_interface_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_instance_network_interfaces table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "network_interface_id",
								Description: "The ID of the network interface.",
								Type:        schema.TypeString,
								Resolver:    schema.ParentPathResolver("NetworkInterfaceId"),
							},
							{
								Name:        "group_id",
								Description: "The ID of the security group.",
								Type:        schema.TypeString,
							},
							{
								Name:        "group_name",
								Description: "The name of the security group.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:          "aws_ec2_instance_network_interface_ipv6_addresses",
						Description:   "Describes an IPv6 address.",
						Resolver:      fetchEc2InstanceNetworkInterfaceIpv6Addresses,
						IgnoreInTests: true,
						Columns: []schema.Column{
							{
								Name:        "instance_network_interface_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_instance_network_interfaces table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:        "ipv6_address",
								Description: "The IPv6 address.",
								Type:        schema.TypeString,
							},
						},
					},
					{
						Name:        "aws_ec2_instance_network_interface_private_ip_addresses",
						Description: "Describes a private IPv4 address.",
						Resolver:    fetchEc2InstanceNetworkInterfacePrivateIpAddresses,
						Columns: []schema.Column{
							{
								Name:        "instance_network_interface_cq_id",
								Description: "Unique CloudQuery ID of aws_ec2_instance_network_interfaces table (FK)",
								Type:        schema.TypeUUID,
								Resolver:    schema.ParentIdResolver,
							},
							{
								Name:          "association_carrier_ip",
								Description:   "The carrier IP address associated with the network interface.",
								Type:          schema.TypeString,
								Resolver:      schema.PathResolver("Association.CarrierIp"),
								IgnoreInTests: true,
							},
							{
								Name:          "association_ip_owner_id",
								Description:   "The ID of the owner of the Elastic IP address.",
								Type:          schema.TypeString,
								Resolver:      schema.PathResolver("Association.IpOwnerId"),
								IgnoreInTests: true,
							},
							{
								Name:        "association_public_dns_name",
								Description: "The public DNS name.",
								Type:        schema.TypeString,
								Resolver:    schema.PathResolver("Association.PublicDnsName"),
							},
							{
								Name:          "association_public_ip",
								Description:   "The public IP address or Elastic IP address bound to the network interface.",
								Type:          schema.TypeString,
								Resolver:      schema.PathResolver("Association.PublicIp"),
								IgnoreInTests: true,
							},
							{
								Name:        "is_primary",
								Description: "Indicates whether this IPv4 address is the primary private IP address of the network interface.",
								Type:        schema.TypeBool,
								Resolver:    schema.PathResolver("Primary"),
							},
							{
								Name:          "private_dns_name",
								Description:   "The private IPv4 DNS name.",
								Type:          schema.TypeString,
								IgnoreInTests: true,
							},
							{
								Name:        "private_ip_address",
								Description: "The private IPv4 address of the network interface.",
								Type:        schema.TypeString,
							},
						},
					},
				},
			},
			{
				Name:          "aws_ec2_instance_product_codes",
				Description:   "Describes a product code.",
				Resolver:      fetchEc2InstanceProductCodes,
				Options:       schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "product_code_id"}},
				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "product_code_id",
						Description: "The product code.",
						Type:        schema.TypeString,
					},
					{
						Name:        "product_code_type",
						Description: "The type of product code.",
						Type:        schema.TypeString,
					},
				},
			},
			{
				Name:        "aws_ec2_instance_security_groups",
				Description: "Describes a security group.",
				Resolver:    fetchEc2InstanceSecurityGroups,
				Options:     schema.TableCreationOptions{PrimaryKeys: []string{"instance_cq_id", "group_id"}},
				Columns: []schema.Column{
					{
						Name:        "instance_cq_id",
						Description: "Unique CloudQuery ID of aws_ec2_instances table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "group_id",
						Description: "The ID of the security group.",
						Type:        schema.TypeString,
					},
					{
						Name:        "group_name",
						Description: "The name of the security group.",
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

func fetchEc2Instances(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().EC2

	response, err := svc.DescribeInstances(ctx, &ec2.DescribeInstancesInput{}, func(o *ec2.Options) {
		o.Region = c.Region
	})
	if err != nil {
		return diag.WrapError(err)
	}

	for _, reservation := range response.Reservations {
		res <- reservation.Instances
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
func resolveEc2InstancesLicenses(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instance := resource.Item.(types.Instance)
	licenses := make([]string, len(instance.Licenses))
	for i, l := range instance.Licenses {
		licenses[i] = *l.LicenseConfigurationArn
	}
	return resource.Set(c.Name, licenses)
}
func resolveEc2InstancesTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.Instance)
	tags := map[string]*string{}
	for _, t := range r.Tags {
		tags[*t.Key] = t.Value
	}
	return resource.Set("tags", tags)
}
func fetchEc2InstanceBlockDeviceMappings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	res <- instance.BlockDeviceMappings
	return nil
}
func fetchEc2InstanceElasticGpuAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	res <- instance.ElasticGpuAssociations
	return nil
}
func fetchEc2InstanceElasticInferenceAcceleratorAssociations(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	res <- instance.ElasticInferenceAcceleratorAssociations
	return nil
}
func fetchEc2InstanceNetworkInterfaces(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	res <- instance.NetworkInterfaces
	return nil
}
func resolveEc2InstanceNetworkInterfacesIpv4Prefixes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instanceNetworkInterface := resource.Item.(types.InstanceNetworkInterface)
	ips := make([]string, 0, len(instanceNetworkInterface.Ipv4Prefixes))
	for _, p := range instanceNetworkInterface.Ipv4Prefixes {
		ips = append(ips, *p.Ipv4Prefix)
	}
	return resource.Set(c.Name, ips)
}
func resolveEc2InstanceNetworkInterfacesIpv6Prefixes(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	instanceNetworkInterface := resource.Item.(types.InstanceNetworkInterface)
	ips := make([]string, 0, len(instanceNetworkInterface.Ipv6Prefixes))
	for _, p := range instanceNetworkInterface.Ipv6Prefixes {
		ips = append(ips, *p.Ipv6Prefix)
	}
	return resource.Set(c.Name, ips)
}
func fetchEc2InstanceNetworkInterfaceGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instanceNetworkInterface := parent.Item.(types.InstanceNetworkInterface)
	res <- instanceNetworkInterface.Groups
	return nil
}
func fetchEc2InstanceNetworkInterfaceIpv6Addresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instanceNetworkInterface := parent.Item.(types.InstanceNetworkInterface)
	res <- instanceNetworkInterface.Ipv6Addresses
	return nil
}
func fetchEc2InstanceNetworkInterfacePrivateIpAddresses(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instanceNetworkInterface := parent.Item.(types.InstanceNetworkInterface)
	res <- instanceNetworkInterface.PrivateIpAddresses
	return nil
}
func fetchEc2InstanceProductCodes(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	res <- instance.ProductCodes
	return nil
}
func fetchEc2InstanceSecurityGroups(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	instance := parent.Item.(types.Instance)
	res <- instance.SecurityGroups
	return nil
}
