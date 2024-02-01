# Table: aws_organizations_delegated_services

This table shows data for Organizations Delegated Services.

https://docs.aws.amazon.com/organizations/latest/APIReference/API_DelegatedService.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **service_principal**).
## Relations

This table depends on [aws_organizations_delegated_administrators](aws_organizations_delegated_administrators.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|delegation_enabled_date|`timestamp[us, tz=UTC]`|
|service_principal|`utf8`|