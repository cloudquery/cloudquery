# Table: azure_peering_service_providers

https://learn.microsoft.com/en-us/rest/api/peering/peering-service-providers/list?tabs=HTTP#peeringserviceprovider

The composite primary key for this table is (**subscription_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|properties|JSON|
|id|String|
|name (PK)|String|
|type|String|