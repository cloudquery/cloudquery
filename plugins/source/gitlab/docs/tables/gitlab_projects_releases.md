# Table: gitlab_projects_releases



The primary key for this table is **_cq_id**.

## Relations
This table depends on [gitlab_projects](gitlab_projects.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|base_url|String|
|tag_name|String|
|name|String|
|description|String|
|description_html|String|
|created_at|Timestamp|
|released_at|Timestamp|
|author|JSON|
|commit|JSON|
|upcoming_release|Bool|
|commit_path|String|
|tag_path|String|
|assets|JSON|