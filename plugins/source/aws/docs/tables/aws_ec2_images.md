# Table: aws_ec2_images

This table shows data for Amazon Elastic Compute Cloud (EC2) Images.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Image.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).
## Relations

The following tables depend on aws_ec2_images:
  - [aws_ec2_image_last_launched_times](aws_ec2_image_last_launched_times.md)
  - [aws_ec2_image_launch_permissions](aws_ec2_image_launch_permissions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|architecture|`utf8`|
|block_device_mappings|`json`|
|boot_mode|`utf8`|
|creation_date|`timestamp[us, tz=UTC]`|
|deprecation_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|ena_support|`bool`|
|hypervisor|`utf8`|
|image_id|`utf8`|
|image_location|`utf8`|
|image_owner_alias|`utf8`|
|image_type|`utf8`|
|imds_support|`utf8`|
|kernel_id|`utf8`|
|name|`utf8`|
|owner_id|`utf8`|
|platform|`utf8`|
|platform_details|`utf8`|
|product_codes|`json`|
|public|`bool`|
|ramdisk_id|`utf8`|
|root_device_name|`utf8`|
|root_device_type|`utf8`|
|source_instance_id|`utf8`|
|sriov_net_support|`utf8`|
|state|`utf8`|
|state_reason|`json`|
|tpm_support|`utf8`|
|usage_operation|`utf8`|
|virtualization_type|`utf8`|