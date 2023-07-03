# Table: gcp_baremetalsolution_volume_luns

This table shows data for GCP Bare Metal Solution Volume Luns.

https://cloud.google.com/bare-metal/docs/reference/rest/v2/projects.locations.volumes.luns#Lun

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_baremetalsolution_volumes](gcp_baremetalsolution_volumes).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|state|`utf8`|
|size_gb|`int64`|
|multiprotocol_type|`utf8`|
|storage_volume|`utf8`|
|shareable|`bool`|
|boot_lun|`bool`|
|storage_type|`utf8`|
|wwid|`utf8`|