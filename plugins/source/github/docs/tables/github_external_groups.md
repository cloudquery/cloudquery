# Table: github_external_groups



The composite primary key for this table is (**org**, **group_id**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|org (PK)|String|
|group_id (PK)|Int|
|updated_at|Timestamp|
|group_name|String|
|teams|JSON|
|members|JSON|