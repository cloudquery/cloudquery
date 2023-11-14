# Table: azure_resources_resources

This table shows data for Azure Resources Resources.

https://learn.microsoft.com/en-us/rest/api/resources/resources/list#genericresourceexpanded

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|extended_location|`json`|
|identity|`json`|
|kind|`utf8`|
|location|`utf8`|
|managed_by|`utf8`|
|plan|`json`|
|sku|`json`|
|tags|`json`|
|changed_time|`timestamp[us, tz=UTC]`|
|created_time|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|name|`utf8`|
|provisioning_state|`utf8`|
|type|`utf8`|