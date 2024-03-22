# Table: aws_rds_option_groups

This table shows data for Amazon Relational Database Service (RDS) Option Groups.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_OptionGroup.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|allows_vpc_and_non_vpc_instance_memberships|`bool`|
|copy_timestamp|`timestamp[us, tz=UTC]`|
|engine_name|`utf8`|
|major_engine_version|`utf8`|
|option_group_arn|`utf8`|
|option_group_description|`utf8`|
|option_group_name|`utf8`|
|options|`json`|
|source_account_id|`utf8`|
|source_option_group|`utf8`|
|vpc_id|`utf8`|