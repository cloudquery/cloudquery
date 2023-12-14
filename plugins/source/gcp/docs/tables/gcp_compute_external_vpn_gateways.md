# Table: gcp_compute_external_vpn_gateways

This table shows data for GCP Compute External VPN Gateways.

https://cloud.google.com/compute/docs/reference/rest/v1/externalVpnGateways#ExternalVpnGateway

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
|interfaces|`json`|
|kind|`utf8`|
|label_fingerprint|`utf8`|
|labels|`json`|
|name|`utf8`|
|redundancy_type|`utf8`|
|self_link (PK)|`utf8`|