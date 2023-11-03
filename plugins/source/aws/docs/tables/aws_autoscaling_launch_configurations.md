# Table: aws_autoscaling_launch_configurations

This table shows data for Auto Scaling Launch Configurations.

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LaunchConfiguration.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|
|image_id|`utf8`|
|instance_type|`utf8`|
|launch_configuration_name|`utf8`|
|associate_public_ip_address|`bool`|
|block_device_mappings|`json`|
|classic_link_vpc_id|`utf8`|
|classic_link_vpc_security_groups|`list<item: utf8, nullable>`|
|ebs_optimized|`bool`|
|iam_instance_profile|`utf8`|
|instance_monitoring|`json`|
|kernel_id|`utf8`|
|key_name|`utf8`|
|launch_configuration_arn|`utf8`|
|metadata_options|`json`|
|placement_tenancy|`utf8`|
|ramdisk_id|`utf8`|
|security_groups|`list<item: utf8, nullable>`|
|spot_price|`utf8`|
|user_data|`utf8`|