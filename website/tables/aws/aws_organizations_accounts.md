# Table: aws_organizations_accounts

This table shows data for Organizations Accounts.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Account.html
The 'request_account_id' column is added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **arn**).

## Relations

The following tables depend on aws_organizations_accounts:
  - [aws_organizations_delegated_services](aws_organizations_delegated_services)

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
|email|String|
|id|String|
|joined_method|String|
|joined_timestamp|Timestamp|
|name|String|
|status|String|