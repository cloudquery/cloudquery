# Table: github_external_groups


The composite primary key for this table is (**org**, **group_id**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|org (PK)|String|
|group_id (PK)|Int|
|updated_at|Timestamp|
|group_name|String|
|teams|JSON|
|members|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|