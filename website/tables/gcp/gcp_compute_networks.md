# Table: gcp_compute_networks

This table shows data for GCP Compute Networks.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|ipv4_range|`utf8`|
|auto_create_subnetworks|`bool`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|enable_ula_internal_ipv6|`bool`|
|firewall_policy|`utf8`|
|gateway_ipv4|`utf8`|
|id|`int64`|
|internal_ipv6_range|`utf8`|
|kind|`utf8`|
|mtu|`int64`|
|name|`utf8`|
|network_firewall_policy_enforcement_order|`utf8`|
|peerings|`json`|
|routing_config|`json`|
|self_link (PK)|`utf8`|
|self_link_with_id|`utf8`|
|subnetworks|`list<item: utf8, nullable>`|