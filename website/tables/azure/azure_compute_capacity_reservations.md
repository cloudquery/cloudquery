# Table: azure_compute_capacity_reservations

This table shows data for Azure Compute Capacity Reservations.

https://learn.microsoft.com/en-us/rest/api/compute/capacity-reservations/list-by-capacity-reservation-group?tabs=HTTP#capacityreservation

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_capacity_reservation_groups](azure_compute_capacity_reservation_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|sku|JSON|
|properties|JSON|
|tags|JSON|
|zones|StringArray|
|id (PK)|String|
|name|String|
|type|String|