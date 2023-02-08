# Table: gcp_compute_routers

https://cloud.google.com/compute/docs/reference/rest/v1/routers/list#response-body

The primary key for this table is **id**.

## Relations

The following tables depend on gcp_compute_routers:
  - [gcp_compute_router_nat_mapping_infos](gcp_compute_router_nat_mapping_infos.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|bgp|JSON|
|bgp_peers|JSON|
|creation_timestamp|String|
|description|String|
|encrypted_interconnect_router|Bool|
|id (PK)|Int|
|interfaces|JSON|
|kind|String|
|md5_authentication_keys|JSON|
|name|String|
|nats|JSON|
|network|String|
|region|String|
|self_link|String|