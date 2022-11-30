# Table: gitlab_groups_group_members



The primary key for this table is **id**.

## Relations
This table depends on [gitlab_groups_groups](gitlab_groups_groups.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|username|String|
|name|String|
|state|String|
|avatar_url|String|
|web_url|String|
|created_at|Timestamp|
|expires_at|JSON|
|access_level|Int|
|group_saml_identity|JSON|