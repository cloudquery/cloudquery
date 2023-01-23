# Table: gcp_baremetalsolution_volume_luns

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.volumes.luns#Lun

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_baremetalsolution_volumes](gcp_baremetalsolution_volumes.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|id|String|
|state|String|
|size_gb|Int|
|multiprotocol_type|String|
|storage_volume|String|
|shareable|Bool|
|boot_lun|Bool|
|storage_type|String|
|wwid|String|