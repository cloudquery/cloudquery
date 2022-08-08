
# Table: aws_ec2_instance_types
Describes the instance type.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|auto_recovery_supported|boolean|Indicates whether auto recovery is supported.|
|bare_metal|boolean|Indicates whether the instance is a bare metal instance type.|
|burstable_performance_supported|boolean|Indicates whether the instance type is a burstable performance instance type.|
|current_generation|boolean|Indicates whether the instance type is current generation.|
|dedicated_hosts_supported|boolean|Indicates whether Dedicated Hosts are supported on the instance type.|
|ebs_info_ebs_optimized_info_baseline_bandwidth_in_mbps|bigint|The baseline bandwidth performance for an EBS-optimized instance type, in Mbps.|
|ebs_info_ebs_optimized_info_baseline_iops|bigint|The baseline input/output storage operations per seconds for an EBS-optimized instance type.|
|ebs_info_ebs_optimized_info_baseline_throughput_in_mb_ps|float|The baseline throughput performance for an EBS-optimized instance type, in MB/s.|
|ebs_info_ebs_optimized_info_maximum_bandwidth_in_mbps|bigint|The maximum bandwidth performance for an EBS-optimized instance type, in Mbps.|
|ebs_info_ebs_optimized_info_maximum_iops|bigint|The maximum input/output storage operations per second for an EBS-optimized instance type.|
|ebs_info_ebs_optimized_info_maximum_throughput_in_mb_ps|float|The maximum throughput performance for an EBS-optimized instance type, in MB/s.|
|ebs_info_ebs_optimized_support|text|Indicates whether the instance type is Amazon EBS-optimized|
|ebs_info_encryption_support|text|Indicates whether Amazon EBS encryption is supported.|
|ebs_info_nvme_support|text|Indicates whether non-volatile memory express (NVMe) is supported.|
|fpga_info_total_fpga_memory_in_mi_b|bigint|The total memory of all FPGA accelerators for the instance type.|
|free_tier_eligible|boolean|Indicates whether the instance type is eligible for the free tier.|
|gpu_info_total_gpu_memory_in_mi_b|bigint|The total size of the memory for the GPU accelerators for the instance type, in MiB.|
|hibernation_supported|boolean|Indicates whether On-Demand hibernation is supported.|
|hypervisor|text|The hypervisor for the instance type.|
|instance_storage_info_encryption_support|text|Indicates whether data is encrypted at rest.|
|instance_storage_info_nvme_support|text|Indicates whether non-volatile memory express (NVMe) is supported.|
|instance_storage_info_total_size_in_gb|bigint|The total size of the disks, in GB.|
|instance_storage_supported|boolean|Indicates whether instance storage is supported.|
|instance_type|text|The instance type|
|memory_info_size_in_mi_b|bigint|The size of the memory, in MiB.|
|network_info_default_network_card_index|bigint|The index of the default network card, starting at 0.|
|network_info_efa_info_maximum_efa_interfaces|bigint|The maximum number of Elastic Fabric Adapters for the instance type.|
|network_info_efa_supported|boolean|Indicates whether Elastic Fabric Adapter (EFA) is supported.|
|network_info_ena_support|text|Indicates whether Elastic Network Adapter (ENA) is supported.|
|network_info_encryption_in_transit_supported|boolean|Indicates whether the instance type automatically encrypts in-transit traffic between instances.|
|network_info_ipv4_addresses_per_interface|bigint|The maximum number of IPv4 addresses per network interface.|
|network_info_ipv6_addresses_per_interface|bigint|The maximum number of IPv6 addresses per network interface.|
|network_info_ipv6_supported|boolean|Indicates whether IPv6 is supported.|
|network_info_maximum_network_cards|bigint|The maximum number of physical network cards that can be allocated to the instance.|
|network_info_maximum_network_interfaces|bigint|The maximum number of network interfaces for the instance type.|
|network_info_network_performance|text|The network performance.|
|placement_group_info_supported_strategies|text[]|The supported placement group types.|
|processor_info_supported_architectures|text[]|The architectures supported by the instance type.|
|processor_info_sustained_clock_speed_in_ghz|float|The speed of the processor, in GHz.|
|supported_boot_modes|text[]|The supported boot modes|
|supported_root_device_types|text[]|The supported root device types.|
|supported_usage_classes|text[]|Indicates whether the instance type is offered for spot or On-Demand.|
|supported_virtualization_types|text[]|The supported virtualization types.|
|v_cpu_info_default_cores|bigint|The default number of cores for the instance type.|
|v_cpu_info_default_threads_per_core|bigint|The default number of threads per core for the instance type.|
|v_cpu_info_default_v_cpus|bigint|The default number of vCPUs for the instance type.|
|v_cpu_info_valid_cores|integer[]|The valid number of cores that can be configured for the instance type.|
|v_cpu_info_valid_threads_per_core|integer[]|The valid number of threads per core that can be configured for the instance type.|
