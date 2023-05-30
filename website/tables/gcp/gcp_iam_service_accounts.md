# Table: gcp_iam_service_accounts

This table shows data for GCP IAM Service Accounts.

https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts#ServiceAccount

The composite primary key for this table is (**unique_id**, **name**).

## Relations

The following tables depend on gcp_iam_service_accounts:
  - [gcp_iam_service_account_keys](gcp_iam_service_account_keys)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|unique_id (PK)|`utf8`|
|name (PK)|`utf8`|
|email|`utf8`|
|display_name|`utf8`|
|etag|`binary`|
|description|`utf8`|
|oauth2_client_id|`utf8`|
|disabled|`bool`|