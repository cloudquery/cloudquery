# Table: gcp_compute_routes

This table shows data for GCP Compute Routes.

https://cloud.google.com/compute/docs/reference/rest/v1/routes/list#response-body

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|as_paths|`json`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|dest_range|`utf8`|
|id|`int64`|
|kind|`utf8`|
|name|`utf8`|
|network|`utf8`|
|next_hop_gateway|`utf8`|
|next_hop_hub|`utf8`|
|next_hop_ilb|`utf8`|
|next_hop_instance|`utf8`|
|next_hop_ip|`utf8`|
|next_hop_network|`utf8`|
|next_hop_peering|`utf8`|
|next_hop_vpn_tunnel|`utf8`|
|priority|`int64`|
|route_status|`utf8`|
|route_type|`utf8`|
|self_link (PK)|`utf8`|
|tags|`list<item: utf8, nullable>`|
|warnings|`json`|