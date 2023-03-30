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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id|String|
|location|String|
|properties|JSON|
|tags|JSON|
|zones|StringArray|
|id (PK)|String|
|name|String|
|type|String|