# Table: azure_consumption_billing_profile_reservation_transactions

This table shows data for Azure Consumption Billing Profile Reservation Transactions.

https://learn.microsoft.com/en-us/rest/api/consumption/reservation-transactions/list?tabs=HTTP#reservationtransaction

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|tags|`list<item: utf8, nullable>`|
|type|`utf8`|