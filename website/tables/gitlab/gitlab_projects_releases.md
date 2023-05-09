# Table: gitlab_projects_releases

This table shows data for Gitlab Projects Releases.

The composite primary key for this table is (**base_url**, **project_id**, **created_at**).

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
|tag_name|String|
|name|String|
|description|String|
|description_html|String|
|created_at (PK)|Timestamp|
|released_at|Timestamp|
|author|JSON|
|commit|JSON|
|upcoming_release|Bool|
|commit_path|String|
|tag_path|String|
|assets|JSON|
|_links|JSON|