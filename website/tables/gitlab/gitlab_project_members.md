# Table: gitlab_project_members

This table shows data for Gitlab Project Members.

The composite primary key for this table is (**base_url**, **project_id**, **id**).

## Relations

This table depends on [gitlab_projects](gitlab_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|base_url (PK)|String|
|project_id (PK)|Int|
|id (PK)|Int|
|username|String|
|email|String|
|name|String|
|state|String|
|created_at|Timestamp|
|expires_at|Timestamp|
|access_level|Int|
|web_url|String|
|avatar_url|String|