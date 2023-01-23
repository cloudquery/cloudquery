# Table: gitlab_group_members

The composite primary key for this table is (**base_url**, **group_id**, **id**).

## Relations

This table depends on [gitlab_groups](gitlab_groups.md).

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
|expires_at|Timestamp|
|username|String|
|name|String|
|state|String|
|avatar_url|String|
|web_url|String|
|created_at|Timestamp|
|access_level|Int|
|group_saml_identity|JSON|