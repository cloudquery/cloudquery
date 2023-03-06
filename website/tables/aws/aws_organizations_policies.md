# Table: aws_organizations_policies

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Policy.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|content|JSON|
|arn (PK)|String|
|aws_managed|Bool|
|description|String|
|id|String|
|name|String|
|type|String|