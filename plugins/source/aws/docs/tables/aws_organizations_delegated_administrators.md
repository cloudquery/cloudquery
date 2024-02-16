# Table: aws_organizations_delegated_administrators

This table shows data for Organizations Delegated Administrators.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_DelegatedAdministrator.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **arn**).
## Relations

The following tables depend on aws_organizations_delegated_administrators:
  - [aws_organizations_delegated_services](aws_organizations_delegated_services.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn|`utf8`|
|delegation_enabled_date|`timestamp[us, tz=UTC]`|
|email|`utf8`|
|id|`utf8`|
|joined_method|`utf8`|
|joined_timestamp|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|status|`utf8`|