# Table: azure_consumption_billing_account_events

This table shows data for Azure Consumption Billing Account Events.

https://learn.microsoft.com/en-us/rest/api/consumption/events/list-by-billing-account?tabs=HTTP#eventsummary

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|e_tag|`utf8`|
|properties|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|