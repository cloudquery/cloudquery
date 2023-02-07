# Table: gcp_compute_router_nat_mapping_infos

https://cloud.google.com/compute/docs/reference/rest/v1/routers/getNatMappingInfo#response-body

The primary key for this table is **_cq_id**.

## Relations

This table depends on [gcp_compute_routers](gcp_compute_routers.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|instance_name|String|
|interface_nat_mappings|JSON|