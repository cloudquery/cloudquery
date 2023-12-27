# Table: azure_network_azure_firewall_fqdn_tags

This table shows data for Azure Network Azure Firewall FQDN Tags.

https://learn.microsoft.com/en-us/rest/api/firewall/azure-firewall-fqdn-tags/list-all?tabs=HTTP#azurefirewallfqdntag

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id (PK)|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|