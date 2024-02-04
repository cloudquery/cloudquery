# Table: hubspot_crm_owners

This table shows data for HubSpot CRM Owners.

https://developers.hubspot.com/docs/api/crm/owners

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|first_name|`utf8`|
|last_name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|archived|`bool`|
|teams|`json`|
|user_id|`int64`|
|email|`utf8`|
|updated_at|`timestamp[us, tz=UTC]`|