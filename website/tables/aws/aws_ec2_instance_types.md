# Table: aws_ec2_instance_types

This table shows data for Amazon Elastic Compute Cloud (EC2) Instance Types.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceTypeInfo.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|auto_recovery_supported|`bool`|
|bare_metal|`bool`|
|burstable_performance_supported|`bool`|
|current_generation|`bool`|
|dedicated_hosts_supported|`bool`|
|ebs_info|`json`|
|fpga_info|`json`|
|free_tier_eligible|`bool`|
|gpu_info|`json`|
|hibernation_supported|`bool`|
|hypervisor|`utf8`|
|inference_accelerator_info|`json`|
|instance_storage_info|`json`|
|instance_storage_supported|`bool`|
|instance_type|`utf8`|
|memory_info|`json`|
|network_info|`json`|
|placement_group_info|`json`|
|processor_info|`json`|
|supported_boot_modes|`list<item: utf8, nullable>`|
|supported_root_device_types|`list<item: utf8, nullable>`|
|supported_usage_classes|`list<item: utf8, nullable>`|
|supported_virtualization_types|`list<item: utf8, nullable>`|
|v_cpu_info|`json`|