# Table: aws_organizations_delegated_services

https://docs.aws.amazon.com/organizations/latest/APIReference/API_DelegatedService.html

The composite primary key for this table is (**account_id**, **service_principal**).

## Relations

This table depends on [aws_organizations_accounts](aws_organizations_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|delegation_enabled_date|Timestamp|
|service_principal (PK)|String|