# Table: aws_autoscaling_launch_configurations

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LaunchConfiguration.html

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
|created_time|Timestamp|
|image_id|String|
|instance_type|String|
|launch_configuration_name|String|
|associate_public_ip_address|Bool|
|block_device_mappings|JSON|
|classic_link_vpc_id|String|
|classic_link_vpc_security_groups|StringArray|
|ebs_optimized|Bool|
|iam_instance_profile|String|
|instance_monitoring|JSON|
|kernel_id|String|
|key_name|String|
|launch_configuration_arn|String|
|metadata_options|JSON|
|placement_tenancy|String|
|ramdisk_id|String|
|security_groups|StringArray|
|spot_price|String|
|user_data|String|