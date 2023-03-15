# Table: azure_network_application_gateways

This table shows data for Azure Network Application Gateways.

https://learn.microsoft.com/en-us/rest/api/application-gateway/application-gateways/list?tabs=HTTP#applicationgateway

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|id (PK)|String|
|identity|JSON|
|location|String|
|properties|JSON|
|tags|JSON|
|zones|StringArray|
|etag|String|
|name|String|
|type|String|