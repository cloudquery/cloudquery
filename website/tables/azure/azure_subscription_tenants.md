# Table: azure_subscription_tenants

This table shows data for Azure Subscription Tenants.

https://learn.microsoft.com/en-us/rest/api/resources/tenants/list?tabs=HTTP#tenantiddescription

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|id (PK)|utf8|
|tenant_id|utf8|