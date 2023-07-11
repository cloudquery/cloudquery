# Table: aws_organizations_policies

This table shows data for Organizations Policies.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Policy.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|content|`json`|
|arn (PK)|`utf8`|
|aws_managed|`bool`|
|description|`utf8`|
|id|`utf8`|
|name|`utf8`|
|type|`utf8`|