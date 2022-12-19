# Table: gcp_iam_service_accounts

https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts#ServiceAccount

The primary key for this table is **unique_id**.

## Relations

The following tables depend on gcp_iam_service_accounts:
  - [gcp_iam_service_account_keys](gcp_iam_service_account_keys.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|description|String|
|disabled|Bool|
|display_name|String|
|email|String|
|etag|String|
|name|String|
|oauth2_client_id|String|
|unique_id (PK)|String|