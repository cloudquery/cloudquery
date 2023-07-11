# Table: azure_compute_capacity_reservation_groups

This table shows data for Azure Compute Capacity Reservation Groups.

https://learn.microsoft.com/en-us/rest/api/compute/capacity-reservation-groups/list-by-resource-group?tabs=HTTP#capacityreservationgroup

The primary key for this table is **id**.

## Relations

The following tables depend on azure_compute_capacity_reservation_groups:
  - [azure_compute_capacity_reservations](azure_compute_capacity_reservations)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|zones|`list<item: utf8, nullable>`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|