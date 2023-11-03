# Table: azure_network_vpn_gateways

This table shows data for Azure Network Virtual Private Network (VPN) Gateways.

https://learn.microsoft.com/en-us/rest/api/virtualwan/vpn-gateways/list?tabs=HTTP#vpngateway

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|subscription_id|`utf8`|
|id (PK)|`utf8`|
|location|`utf8`|
|properties|`json`|
|tags|`json`|
|etag|`utf8`|
|name|`utf8`|
|type|`utf8`|