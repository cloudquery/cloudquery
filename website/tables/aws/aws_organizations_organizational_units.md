# Table: aws_organizations_organizational_units

This table shows data for Organizations Organizational Units.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_OrganizationalUnit.html

The composite primary key for this table is (**account_id**, **arn**).

## Relations

The following tables depend on aws_organizations_organizational_units:
  - [aws_organizations_organizational_unit_accounts](aws_organizations_organizational_unit_accounts)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|parent_id|`utf8`|
|arn (PK)|`utf8`|
|id|`utf8`|
|name|`utf8`|