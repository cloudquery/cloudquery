# Table: aws_backup_region_settings

This table shows data for Backup Region Settings.

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeRegionSettings.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|resource_type_management_preference|`json`|
|resource_type_opt_in_preference|`json`|