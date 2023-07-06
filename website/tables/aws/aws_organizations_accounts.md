# Table: aws_organizations_accounts

This table shows data for Organizations Accounts.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_Account.html
The 'request_account_id' column is added to show from where the request was made.

The composite primary key for this table is (**request_account_id**, **arn**).

## Relations

The following tables depend on aws_organizations_accounts:
  - [aws_organizations_account_parents](aws_organizations_account_parents)
  - [aws_organizations_delegated_services](aws_organizations_delegated_services)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|email|`utf8`|
|id|`utf8`|
|joined_method|`utf8`|
|joined_timestamp|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|status|`utf8`|