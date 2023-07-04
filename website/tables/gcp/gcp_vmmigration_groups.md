# Table: gcp_vmmigration_groups

This table shows data for GCP VM Migration Groups.

https://cloud.google.com/migrate/virtual-machines/docs/5.0/reference/rest/v1/projects.locations.groups

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|display_name|`utf8`|