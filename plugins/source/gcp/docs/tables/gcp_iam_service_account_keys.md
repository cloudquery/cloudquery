# Table: gcp_iam_service_account_keys



The primary key for this table is **service_account_unique_id**.

## Relations
This table depends on [gcp_iam_service_accounts](gcp_iam_service_accounts.md).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|service_account_unique_id (PK)|String|
|disabled|Bool|
|key_algorithm|String|
|key_origin|String|
|key_type|String|
|name|String|
|private_key_type|String|
|public_key_data|String|
|valid_after_time|String|
|valid_before_time|String|