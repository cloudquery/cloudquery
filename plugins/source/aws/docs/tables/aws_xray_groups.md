# Table: aws_xray_groups

This table shows data for AWS X-Ray Groups.

https://docs.aws.amazon.com/xray/latest/api/API_Group.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|filter_expression|`utf8`|
|group_arn|`utf8`|
|group_name|`utf8`|
|insights_configuration|`json`|