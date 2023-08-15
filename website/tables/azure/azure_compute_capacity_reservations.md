# Table: azure_compute_capacity_reservations

This table shows data for Azure Compute Capacity Reservations.

https://learn.microsoft.com/en-us/rest/api/compute/capacity-reservations/list-by-capacity-reservation-group?tabs=HTTP#capacityreservation

The primary key for this table is **id**.

## Relations

This table depends on [azure_compute_capacity_reservation_groups](azure_compute_capacity_reservation_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|sku|`json`|
|properties|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|