# Table: aws_organizations_roots

This table shows data for Organizations Roots.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Root.html
The 'request_account_id' column is added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|request_account_id (PK)|String|
|tags|JSON|
|arn (PK)|String|
|id|String|
|name|String|
|policy_types|JSON|