# Table: azure_network_azure_firewall_fqdn_tags

This table shows data for Azure Network Azure Firewall FQDN Tags.

https://learn.microsoft.com/en-us/rest/api/firewall/azure-firewall-fqdn-tags/list-all?tabs=HTTP#azurefirewallfqdntag

The composite primary key for this table is (**subscription_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|subscription_id (PK)|String|
|id (PK)|String|
|location|String|
|properties|JSON|
|tags|JSON|
|etag|String|
|name|String|
|type|String|