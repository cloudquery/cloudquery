# Table: gcp_compute_routers

This table shows data for GCP Compute Routers.

https://cloud.google.com/compute/docs/reference/rest/v1/routers/list#response-body

The primary key for this table is **id**.

## Relations

The following tables depend on gcp_compute_routers:
  - [gcp_compute_router_nat_mapping_infos](gcp_compute_router_nat_mapping_infos)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|bgp|`json`|
|bgp_peers|`json`|
|creation_timestamp|`utf8`|
|description|`utf8`|
|encrypted_interconnect_router|`bool`|
|id (PK)|`int64`|
|interfaces|`json`|
|kind|`utf8`|
|md5_authentication_keys|`json`|
|name|`utf8`|
|nats|`json`|
|network|`utf8`|
|region|`utf8`|
|self_link|`utf8`|