# Table: gcp_compute_osconfig_inventories

This table shows data for GCP Compute OS Config Inventories.

https://cloud.google.com/compute/docs/osconfig/rest/v1/projects.locations.instances.inventories#Inventory

The primary key for this table is **name**.

## Relations

This table depends on [gcp_compute_zones](gcp_compute_zones).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|os_info|`json`|
|items|`json`|
|update_time|`timestamp[us, tz=UTC]`|