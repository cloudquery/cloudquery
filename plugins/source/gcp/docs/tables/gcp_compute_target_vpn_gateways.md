# Table: gcp_compute_target_vpn_gateways

This table shows data for GCP Compute Target VPN Gateways.

https://cloud.google.com/compute/docs/reference/rest/v1/targetVpnGateways#TargetVpnGateway

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|forwarding_rules|`list<item: utf8, nullable>`|
|id|`int64`|
|kind|`utf8`|
|label_fingerprint|`utf8`|
|labels|`json`|
|name|`utf8`|
|network|`utf8`|
|region|`utf8`|
|self_link (PK)|`utf8`|
|status|`utf8`|
|tunnels|`list<item: utf8, nullable>`|