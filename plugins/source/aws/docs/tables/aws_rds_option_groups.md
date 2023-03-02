# Table: aws_rds_option_groups

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_OptionGroup.html

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
|allows_vpc_and_non_vpc_instance_memberships|Bool|
|copy_timestamp|Timestamp|
|engine_name|String|
|major_engine_version|String|
|option_group_arn|String|
|option_group_description|String|
|option_group_name|String|
|options|JSON|
|source_account_id|String|
|source_option_group|String|
|vpc_id|String|