# Table: gcp_compute_router_nat_mapping_infos

This table shows data for GCP Compute Router Nat Mapping Infos.

https://cloud.google.com/compute/docs/reference/rest/v1/routers/getNatMappingInfo#response-body

The composite primary key for this table is ().

## Relations

This table depends on [gcp_compute_routers](gcp_compute_routers).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|instance_name|`utf8`|
|interface_nat_mappings|`json`|