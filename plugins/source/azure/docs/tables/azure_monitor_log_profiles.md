# Table: azure_monitor_log_profiles

This table shows data for Azure Monitor Log Profiles.

https://learn.microsoft.com/en-us/rest/api/monitor/log-profiles/list?tabs=HTTP#logprofileresource

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|id (PK)|`utf8`|
|name|`utf8`|
|type|`utf8`|