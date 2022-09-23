# Table: gcp_compute_networks


The primary key for this table is **self_link**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|self_link (PK)|String|
|ipv4_range|String|
|auto_create_subnetworks|Bool|
|creation_timestamp|String|
|description|String|
|enable_ula_internal_ipv6|Bool|
|firewall_policy|String|
|gateway_ipv4|String|
|id|Int|
|internal_ipv6_range|String|
|kind|String|
|mtu|Int|
|name|String|
|network_firewall_policy_enforcement_order|String|
|peerings|JSON|
|routing_config|JSON|
|self_link_with_id|String|
|subnetworks|StringArray|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|