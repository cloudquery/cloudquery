# Table: azure_network_firewall_policies

This table shows data for Azure Network Firewall Policies.

https://learn.microsoft.com/en-us/rest/api/virtualnetwork/firewall-policies/list-all?tabs=HTTP#firewallpolicy

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|identity|`json`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|