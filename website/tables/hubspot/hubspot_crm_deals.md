# Table: hubspot_crm_deals

This table shows data for HubSpot CRM Deals.

https://developers.hubspot.com/docs/api/crm/deals

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|properties|`json`|
|created_at|`timestamp[us, tz=UTC]`|
|updated_at|`timestamp[us, tz=UTC]`|
|archived|`bool`|
|archived_at|`timestamp[us, tz=UTC]`|
|associations|`json`|