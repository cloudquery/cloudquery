# Table: features

This table shows data for Features.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|created_by_user|`json`|
|created_at|`int64`|
|last_updated_by_user|`json`|
|last_updated_at|`int64`|
|kind|`utf8`|
|root_version_id|`utf8`|
|stable_version_id|`utf8`|
|id|`utf8`|
|app_id|`int64`|
|name|`utf8`|
|color|`utf8`|
|group|`json`|
|is_core_event|`bool`|
|valid_through|`int64`|
|dirty|`bool`|
|daily_merge_first|`int64`|
|daily_rollup_first|`int64`|
|page_id|`utf8`|
|event_property_configurations|`json`|
|element_path_rules|`list<item: utf8, nullable>`|
|element_selection_type|`utf8`|
|suggested_match|`utf8`|
|app_wide|`bool`|