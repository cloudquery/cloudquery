# Table: gcp_iam_service_account_keys

This table shows data for GCP IAM Service Account Keys.

https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_iam_service_accounts](gcp_iam_service_accounts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id (PK)|utf8|
|service_account_unique_id|utf8|
|name (PK)|utf8|
|key_algorithm|utf8|
|public_key_data|binary|
|valid_after_time|timestamp[us, tz=UTC]|
|valid_before_time|timestamp[us, tz=UTC]|
|key_origin|utf8|
|key_type|utf8|
|disabled|bool|