# Table: guides

This table shows data for Guides.

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
|app_id|`int64`|
|app_ids|`list<item: int64, nullable>`|
|id|`utf8`|
|name|`utf8`|
|description|`utf8`|
|state|`utf8`|
|email_state|`utf8`|
|launch_method|`utf8`|
|is_multi_step|`bool`|
|is_training|`bool`|
|steps|`json`|
|attributes|`json`|
|audience|`json`|
|audience_ui_hint|`json`|
|authored_language|`utf8`|
|recurrence|`int64`|
|recurrence_eligibility_window|`int64`|
|reset_at|`int64`|
|published_at|`int64`|
|published_ever|`bool`|
|current_first_eligible_to_be_seen_at|`int64`|
|expires_after|`int64`|
|is_top_level|`bool`|
|is_module|`bool`|
|editor_type|`utf8`|
|dependent_metadata|`list<item: utf8, nullable>`|
|shows_after|`int64`|
|polls|`json`|
|translation_states|`json`|