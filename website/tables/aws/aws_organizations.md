# Table: aws_organizations

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Organization.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|arn (PK)|String|
|feature_set|String|
|id|String|
|master_account_arn|String|
|master_account_email|String|
|master_account_id|String|