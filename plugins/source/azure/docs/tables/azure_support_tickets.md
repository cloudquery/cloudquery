# Table: azure_support_tickets

This table shows data for Azure Support Tickets.

https://learn.microsoft.com/en-us/rest/api/support/support-tickets/list?tabs=HTTP#supportticketdetails

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|