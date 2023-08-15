# Table: aws_organizations_delegated_services

This table shows data for Organizations Delegated Services.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_DelegatedService.html

The composite primary key for this table is (**account_id**, **service_principal**).

## Relations

This table depends on [aws_organizations_accounts](aws_organizations_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|delegation_enabled_date|`timestamp[us, tz=UTC]`|
|service_principal (PK)|`utf8`|