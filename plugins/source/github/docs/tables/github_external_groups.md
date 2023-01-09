# Table: github_external_groups

The primary key for this table is **org**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|group_id|Int|
|group_name|String|
|updated_at|Timestamp|
|teams|JSON|
|members|JSON|