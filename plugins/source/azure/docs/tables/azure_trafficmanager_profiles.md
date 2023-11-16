# Table: azure_trafficmanager_profiles

This table shows data for Azure Trafficmanager Profiles.

https://learn.microsoft.com/en-us/rest/api/trafficmanager/profiles/list-by-subscription?tabs=HTTP#profile

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|name|`utf8`|
|properties|`json`|
|tags|`json`|
|type|`utf8`|