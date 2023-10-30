# Table: atlas_datalake_pipelines

This table shows data for Atlas Datalake Pipelines.

The composite primary key for this table is (**_id**, **group_id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|_id (PK)|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|group_id (PK)|`utf8`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|sink|`json`|
|source|`json`|
|state|`utf8`|
|transformations|`json`|