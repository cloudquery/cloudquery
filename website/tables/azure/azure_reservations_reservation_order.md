# Table: azure_reservations_reservation_order

This table shows data for Azure Reservations Reservation Order.

https://learn.microsoft.com/en-us/rest/api/reserved-vm-instances/reservation-order/get?tabs=HTTP#reservationorderresponse

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|etag|`int64`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|