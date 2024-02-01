# Table: aws_organizations_policies

This table shows data for Organizations Policies.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Policy.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|content|`json`|
|arn|`utf8`|
|aws_managed|`bool`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|type|`utf8`|