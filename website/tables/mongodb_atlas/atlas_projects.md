# Table: atlas_projects

This table shows data for Atlas Projects.

The composite primary key for this table is (**id**, **org_id**).

## Relations

The following tables depend on atlas_projects:
  - [atlas_project_settings](atlas_project_settings.md)
  - [atlas_project_users](atlas_project_users.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|cluster_count|`int64`|
|created|`timestamp[us, tz=UTC]`|
|id (PK)|`utf8`|
|links|`json`|
|name|`utf8`|
|org_id (PK)|`utf8`|
|region_usage_restrictions|`utf8`|
|with_default_alerts_settings|`bool`|