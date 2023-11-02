# Table: azure_automation_account

This table shows data for Azure Automation Account.

https://learn.microsoft.com/en-us/rest/api/automation/automation-account/list?tabs=HTTP#automationaccount

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|etag|`utf8`|
|identity|`json`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|system_data|`json`|
|type|`utf8`|