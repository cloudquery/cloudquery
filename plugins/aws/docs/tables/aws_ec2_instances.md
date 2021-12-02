
# Table: aws_ec2_instances
Describes an instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|state_transition_reason_time|timestamp without time zone||
|ami_launch_index|integer|The AMI launch index, which can be used to find this instance in the launch group.|
|architecture|text|The architecture of the image.|
|boot_mode|text|The boot mode of the instance|
|capacity_reservation_id|text|The ID of the Capacity Reservation.|
|cap_reservation_preference|text|Describes the instance's Capacity Reservation preferences|
|cap_reservation_target_capacity_reservation_id|text|The ID of the targeted Capacity Reservation.|
|cap_reservation_target_capacity_reservation_rg_arn|text|The ARN of the targeted Capacity Reservation group.|
|client_token|text|The idempotency token you provided when you launched the instance, if applicable.|
|cpu_options_core_count|integer|The number of CPU cores for the instance.|
|cpu_options_threads_per_core|integer|The number of threads per CPU core.|
|ebs_optimized|boolean|Indicates whether the instance is optimized for Amazon EBS I/O|
|ena_support|boolean|Specifies whether enhanced networking with ENA is enabled.|
|enclave_options_enabled|boolean|If this parameter is set to true, the instance is enabled for Amazon Web Services Nitro Enclaves; otherwise, it is not enabled for Amazon Web Services Nitro Enclaves.|
|hibernation_options_configured|boolean|If this parameter is set to true, your instance is enabled for hibernation; otherwise, it is not enabled for hibernation.|
|hypervisor|text|The hypervisor type of the instance|
|iam_instance_profile_arn|text|The Amazon Resource Name (ARN) of the instance profile.|
|iam_instance_profile_id|text|The ID of the instance profile.|
|image_id|text|The ID of the AMI used to launch the instance.|
|id|text|The ID of the instance.|
|instance_lifecycle|text|Indicates whether this is a Spot Instance or a Scheduled Instance.|
|instance_type|text|The instance type.|
|kernel_id|text|The kernel associated with this instance, if applicable.|
|key_name|text|The name of the key pair, if this instance was launched with an associated key pair.|
|launch_time|timestamp without time zone|The time the instance was launched.|
|licenses|text[]|The license configurations.|
|metadata_options_http_endpoint|text|This parameter enables or disables the HTTP metadata endpoint on your instances. If the parameter is not specified, the default state is enabled|
|metadata_options_http_protocol_ipv6|text|Whether or not the IPv6 endpoint for the instance metadata service is enabled or disabled.|
|metadata_options_http_put_response_hop_limit|integer|The desired HTTP PUT response hop limit for instance metadata requests|
|metadata_options_http_tokens|text|The state of token usage for your instance metadata requests|
|metadata_options_state|text|The state of the metadata option changes|
|monitoring_state|text|Indicates whether detailed monitoring is enabled|
|outpost_arn|text|The Amazon Resource Name (ARN) of the Outpost.|
|placement_affinity|text|The affinity setting for the instance on the Dedicated Host|
|placement_availability_zone|text|The Availability Zone of the instance|
|placement_group_name|text|The name of the placement group the instance is in.|
|placement_host_id|text|The ID of the Dedicated Host on which the instance resides|
|placement_host_resource_group_arn|text|The ARN of the host resource group in which to launch the instances|
|placement_partition_number|integer|The number of the partition the instance is in|
|placement_spread_domain|text|Reserved for future use|
|placement_tenancy|text|The tenancy of the instance (if the instance is running in a VPC)|
|platform|text|The value is Windows for Windows instances; otherwise blank.|
|private_dns_name|text|(IPv4 only) The private DNS hostname name assigned to the instance|
|private_ip_address|text|The private IPv4 address assigned to the instance.|
|public_dns_name|text|(IPv4 only) The public DNS name assigned to the instance|
|public_ip_address|text|The public IPv4 address, or the Carrier IP address assigned to the instance, if applicable|
|ramdisk_id|text|The RAM disk associated with this instance, if applicable.|
|root_device_name|text|The device name of the root device volume (for example, /dev/sda1).|
|root_device_type|text|The root device type used by the AMI|
|source_dest_check|boolean|Indicates whether source/destination checking is enabled.|
|spot_instance_request_id|text|If the request is a Spot Instance request, the ID of the request.|
|sriov_net_support|text|Specifies whether enhanced networking with the Intel 82599 Virtual Function interface is enabled.|
|state_code|integer|The state of the instance as a 16-bit unsigned integer|
|state_name|text|The current state of the instance.|
|state_reason_code|text|The reason code for the state change.|
|state_reason_message|text|The message for the state change.  * Server.InsufficientInstanceCapacity: There was insufficient capacity available to satisfy the launch request.  * Server.InternalError: An internal error caused the instance to terminate during launch.  * Server.ScheduledStop: The instance was stopped due to a scheduled retirement.  * Server.SpotInstanceShutdown: The instance was stopped because the number of Spot requests with a maximum price equal to or higher than the Spot price exceeded available capacity or because of an increase in the Spot price.  * Server.SpotInstanceTermination: The instance was terminated because the number of Spot requests with a maximum price equal to or higher than the Spot price exceeded available capacity or because of an increase in the Spot price.  * Client.InstanceInitiatedShutdown: The instance was shut down using the shutdown -h command from the instance.  * Client.InstanceTerminated: The instance was terminated or rebooted during AMI creation.  * Client.InternalError: A client error caused the instance to terminate during launch.  * Client.InvalidSnapshot.NotFound: The specified snapshot was not found.  * Client.UserInitiatedHibernate: Hibernation was initiated on the instance.  * Client.UserInitiatedShutdown: The instance was shut down using the Amazon EC2 API.  * Client.VolumeLimitExceeded: The limit on the number of EBS volumes or total storage was exceeded|
|state_transition_reason|text|The reason for the most recent state transition|
|subnet_id|text|[EC2-VPC] The ID of the subnet in which the instance is running.|
|tags|jsonb|Any tags assigned to the instance.|
|virtualization_type|text|The virtualization type of the instance.|
|vpc_id|text|The ID of the VPC in which the instance is running.|
