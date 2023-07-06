# Table: aws_organizations_delegated_administrators

This table shows data for Organizations Delegated Administrators.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_DelegatedAdministrator.html

The composite primary key for this table is (**account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|arn (PK)|`utf8`|
|delegation_enabled_date|`timestamp[us, tz=UTC]`|
|email|`utf8`|
|id|`utf8`|
|joined_method|`utf8`|
|joined_timestamp|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|status|`utf8`|