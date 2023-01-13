# Table: gcp_iam_roles

https://cloud.google.com/iam/docs/reference/rest/v1/roles#Role

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|title|String|
|description|String|
|included_permissions|StringArray|
|stage|String|
|etag|ByteArray|
|deleted|Bool|