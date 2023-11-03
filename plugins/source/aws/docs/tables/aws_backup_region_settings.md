# Table: aws_backup_region_settings

This table shows data for Backup Region Settings.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeRegionSettings.html

The composite primary key for this table is (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|resource_type_management_preference|`json`|
|resource_type_opt_in_preference|`json`|
|result_metadata|`json`|