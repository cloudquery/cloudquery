# Table: gcp_secretmanager_secrets

This table shows data for GCP Secret Manager Secrets.

https://cloud.google.com/secret-manager/docs/reference/rest/v1/projects.secrets#Secret

The primary key for this table is **name**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|name (PK)|`utf8`|
|replication|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|topics|`json`|
|etag|`utf8`|
|rotation|`json`|
|version_aliases|`json`|
|annotations|`json`|