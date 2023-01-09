# Table: gcp_iam_deny_policies

https://cloud.google.com/iam/docs/reference/rest/v2beta/policies#Policy

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|name|String|
|uid|String|
|kind|String|
|display_name|String|
|annotations|JSON|
|etag|String|
|create_time|Timestamp|
|update_time|Timestamp|
|delete_time|Timestamp|
|rules|JSON|
|managing_authority|String|