# Table: azure_synapse_workspaces

This table shows data for Azure Synapse Workspaces.

https://learn.microsoft.com/en-us/rest/api/synapse/workspaces/list?tabs=HTTP#workspace

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|identity|`json`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|