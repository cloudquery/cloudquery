# Table: github_repository_branches

This table shows data for Github Repository Branches.

The composite primary key for this table is (**org**, **repository_id**, **name**).

## Relations

This table depends on [github_repositories](github_repositories).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|repository_id (PK)|Int|
|name (PK)|String|
|commit|JSON|
|protected|Bool|
|protection|JSON|
