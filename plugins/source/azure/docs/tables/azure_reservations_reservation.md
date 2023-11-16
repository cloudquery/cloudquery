# Table: azure_reservations_reservation

This table shows data for Azure Reservations Reservation.

https://learn.microsoft.com/en-us/rest/api/reserved-vm-instances/reservation/list-all?tabs=HTTP#reservationresponse

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|etag|`int64`|
|kind|`utf8`|
|location|`utf8`|
|properties|`json`|
|sku|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|