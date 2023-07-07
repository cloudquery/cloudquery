# Table: aws_ec2_images

This table shows data for Amazon Elastic Compute Cloud (EC2) Images.

https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_Image.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Relations

The following tables depend on aws_ec2_images:
  - [aws_ec2_image_last_launched_times](aws_ec2_image_last_launched_times)
  - [aws_ec2_image_launch_permissions](aws_ec2_image_launch_permissions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|architecture|`utf8`|
|block_device_mappings|`json`|
|boot_mode|`utf8`|
|creation_date|`utf8`|
|deprecation_time|`utf8`|
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
|sriov_net_support|`utf8`|
|state|`utf8`|
|state_reason|`json`|
|tpm_support|`utf8`|
|usage_operation|`utf8`|
|virtualization_type|`utf8`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### AMIs should require IMDSv2

```sql
SELECT
  'AMIs should require IMDSv2' AS title,
  account_id,
  arn AS resource_id,
  CASE
  WHEN imds_support IS DISTINCT FROM 'v2.0' THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_ec2_images;
```

### Unused own EC2 image

```sql
SELECT
  'Unused own EC2 image' AS title,
  account_id,
  arn AS resource_id,
  'fail' AS status
FROM
  aws_ec2_images
WHERE
  COALESCE(jsonb_array_length(block_device_mappings), 0) = 0;
```


