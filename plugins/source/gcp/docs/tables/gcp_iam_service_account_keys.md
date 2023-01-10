# Table: gcp_iam_service_account_keys

https://cloud.google.com/iam/docs/reference/rest/v1/projects.serviceAccounts.keys#ServiceAccountKey

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_iam_service_accounts](gcp_iam_service_accounts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|service_account_unique_id|String|
|name (PK)|String|
|key_algorithm|String|
|public_key_data|ByteArray|
|valid_after_time|Timestamp|
|valid_before_time|Timestamp|
|key_origin|String|
|key_type|String|
|disabled|Bool|