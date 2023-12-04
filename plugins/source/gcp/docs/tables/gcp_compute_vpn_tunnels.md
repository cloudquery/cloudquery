# Table: gcp_compute_vpn_tunnels

This table shows data for GCP Compute VPN Tunnels.

https://cloud.google.com/compute/docs/reference/rest/v1/vpnTunnels#VpnTunnel

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|detailed_status|`utf8`|
|id|`int64`|
|ike_version|`int64`|
|kind|`utf8`|
|label_fingerprint|`utf8`|
|labels|`json`|
|local_traffic_selector|`list<item: utf8, nullable>`|
|name|`utf8`|
|peer_external_gateway|`utf8`|
|peer_external_gateway_interface|`int64`|
|peer_gcp_gateway|`utf8`|
|peer_ip|`utf8`|
|region|`utf8`|
|remote_traffic_selector|`list<item: utf8, nullable>`|
|router|`utf8`|
|self_link (PK)|`utf8`|
|shared_secret|`utf8`|
|shared_secret_hash|`utf8`|
|status|`utf8`|
|target_vpn_gateway|`utf8`|
|vpn_gateway|`utf8`|
|vpn_gateway_interface|`int64`|