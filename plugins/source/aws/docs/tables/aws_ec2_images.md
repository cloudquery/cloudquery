# Table: aws_ec2_images

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Image.html

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
|architecture|String|
|block_device_mappings|JSON|
|boot_mode|String|
|creation_date|String|
|deprecation_time|String|
|description|String|
|ena_support|Bool|
|hypervisor|String|
|image_id|String|
|image_location|String|
|image_owner_alias|String|
|image_type|String|
|imds_support|String|
|kernel_id|String|
|name|String|
|owner_id|String|
|platform|String|
|platform_details|String|
|product_codes|JSON|
|public|Bool|
|ramdisk_id|String|
|root_device_name|String|
|root_device_type|String|
|sriov_net_support|String|
|state|String|
|state_reason|JSON|
|tags|JSON|
|tpm_support|String|
|usage_operation|String|
|virtualization_type|String|