# Table: aws_ec2_instance_types

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_InstanceTypeInfo.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|auto_recovery_supported|Bool|
|bare_metal|Bool|
|burstable_performance_supported|Bool|
|current_generation|Bool|
|dedicated_hosts_supported|Bool|
|ebs_info|JSON|
|fpga_info|JSON|
|free_tier_eligible|Bool|
|gpu_info|JSON|
|hibernation_supported|Bool|
|hypervisor|String|
|inference_accelerator_info|JSON|
|instance_storage_info|JSON|
|instance_storage_supported|Bool|
|instance_type|String|
|memory_info|JSON|
|network_info|JSON|
|placement_group_info|JSON|
|processor_info|JSON|
|supported_boot_modes|StringArray|
|supported_root_device_types|StringArray|
|supported_usage_classes|StringArray|
|supported_virtualization_types|StringArray|
|v_cpu_info|JSON|