# Table: aws_organizations_root_accounts

This table shows data for Organizations Root Accounts.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListAccountsForParent.html

The composite primary key for this table is (**request_account_id**, **parent_id**, **arn**).

## Relations

This table depends on [aws_organizations_roots](aws_organizations_roots).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|parent_id (PK)|`utf8`|
|arn (PK)|`utf8`|
|email|`utf8`|
|id|`utf8`|
|joined_method|`utf8`|
|joined_timestamp|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|status|`utf8`|