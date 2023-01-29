# Table: azure_reservations_reservation

https://learn.microsoft.com/en-us/rest/api/reserved-vm-instances/reservation/get?tabs=HTTP#reservationresponse

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|etag|Int|
|kind|String|
|location|String|
|properties|JSON|
|sku|JSON|
|id (PK)|String|
|name|String|
|system_data|JSON|
|type|String|