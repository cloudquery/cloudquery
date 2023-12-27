# Table: gcp_compute_vpn_gateways

This table shows data for GCP Compute VPN Gateways.

https://cloud.google.com/compute/docs/reference/rest/v1/vpnGateways#VpnGateway

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|id|`int64`|
|kind|`utf8`|
|label_fingerprint|`utf8`|
|labels|`json`|
|name|`utf8`|
|network|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|stack_type|`utf8`|
|vpn_interfaces|`json`|