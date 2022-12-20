# Table: gcp_secretmanager_secrets

https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets#Secret

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|name (PK)|String|
|replication|JSON|
|create_time|Timestamp|
|labels|JSON|
|topics|JSON|
|etag|String|
|rotation|JSON|
|version_aliases|JSON|