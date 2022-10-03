# Table: gcp_iam_service_accounts


The primary key for this table is **unique_id**.

## Relations
The following tables depend on `gcp_iam_service_accounts`:
  - [`gcp_iam_service_account_keys`](gcp_iam_service_account_keys.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|unique_id (PK)|String|
|description|String|
|disabled|Bool|
|display_name|String|
|email|String|
|etag|String|
|name|String|
|oauth2_client_id|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|