# Table: azure_logic_workflows

This table shows data for Azure Logic Workflows.

https://learn.microsoft.com/en-us/rest/api/logic/workflows/list-by-subscription?tabs=HTTP#workflow

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|identity|`json`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|