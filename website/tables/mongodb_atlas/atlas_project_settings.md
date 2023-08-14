# Table: atlas_project_settings

This table shows data for Atlas Project Settings.

The primary key for this table is **project_id**.

## Relations

This table depends on [atlas_projects](atlas_projects.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|is_collect_database_specifics_statistics_enabled|`bool`|
|is_data_explorer_enabled|`bool`|
|is_extended_storage_sizes_enabled|`bool`|
|is_performance_advisor_enabled|`bool`|
|is_realtime_performance_panel_enabled|`bool`|
|is_schema_advisor_enabled|`bool`|