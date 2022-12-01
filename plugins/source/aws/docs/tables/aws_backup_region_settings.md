# Table: aws_backup_region_settings

https://docs.aws.amazon.com/aws-backup/latest/devguide/API_DescribeRegionSettings.html

The composite primary key for this table is (**account_id**, **region**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|resource_type_management_preference|JSON|
|resource_type_opt_in_preference|JSON|
|result_metadata|JSON|