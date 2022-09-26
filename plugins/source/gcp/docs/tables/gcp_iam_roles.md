# Table: gcp_iam_roles


The composite primary key for this table is (**project_id**, **name**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id (PK)|String|
|name (PK)|String|
|deleted|Bool|
|description|String|
|etag|String|
|included_permissions|StringArray|
|stage|String|
|title|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|