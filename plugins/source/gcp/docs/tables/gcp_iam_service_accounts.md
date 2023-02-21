# Table: gcp_iam_service_accounts

https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts#ServiceAccount

The composite primary key for this table is (**unique_id**, **name**).

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
|unique_id (PK)|String|
|name (PK)|String|
|email|String|
|display_name|String|
|etag|ByteArray|
|description|String|
|oauth2_client_id|String|
|disabled|Bool|