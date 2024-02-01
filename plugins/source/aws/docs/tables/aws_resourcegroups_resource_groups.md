# Table: aws_resourcegroups_resource_groups

This table shows data for Resourcegroups Resource Groups.

https://docs.aws.amazon.com/ARG/latest/APIReference/API_GetGroupQuery.html

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
|group_arn|`utf8`|
|name|`utf8`|
|description|`utf8`|
|query|`utf8`|
|type|`utf8`|