# Table: gitlab_project_branches

This table shows data for Gitlab Project Branches.

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gitlab_projects](gitlab_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|base_url|String|
|project_id (PK)|Int|
|name (PK)|String|
|commit|JSON|
|protected|Bool|
|merged|Bool|
|default|Bool|
|can_push|Bool|
|developers_can_push|Bool|
|developers_can_merge|Bool|
|web_url|String|