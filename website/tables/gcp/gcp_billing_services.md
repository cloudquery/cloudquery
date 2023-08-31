# Table: gcp_billing_services

This table shows data for GCP Billing Services.

https://cloud.google.com/billing/docs/reference/rest/v1/services/list#Service

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|service_id|`utf8`|
|display_name|`utf8`|
|business_entity_name|`utf8`|