# Table: gitlab_group_billable_members

This table shows data for Gitlab Group Billable Members.

The composite primary key for this table is (**base_url**, **group_id**, **id**).

## Relations

This table depends on [gitlab_groups](gitlab_groups).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|base_url (PK)|String|
|group_id (PK)|Int|
|id (PK)|Int|
|username|String|
|name|String|
|state|String|
|avatar_url|String|
|web_url|String|
|email|String|
|last_activity_on|Timestamp|
|membership_type|String|
|removable|Bool|
|created_at|Timestamp|
|is_last_owner|Bool|
|last_login_at|Timestamp|