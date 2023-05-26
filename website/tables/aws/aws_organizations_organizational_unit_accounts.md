# Table: aws_organizations_organizational_unit_accounts

This table shows data for Organizations Organizational Unit Accounts.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_ListAccountsForParent.html

The composite primary key for this table is (**account_id**, **parent_id**, **arn**).

## Relations

This table depends on [aws_organizations_organizational_units](aws_organizations_organizational_units).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|parent_id (PK)|utf8|
|arn (PK)|utf8|
|email|utf8|
|id|utf8|
|joined_method|utf8|
|joined_timestamp|timestamp[us, tz=UTC]|
|name|utf8|
|status|utf8|