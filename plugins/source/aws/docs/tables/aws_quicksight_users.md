# Table: aws_quicksight_users



The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|arn (PK)|String|
|active|Bool|
|custom_permissions_name|String|
|email|String|
|external_login_federation_provider_type|String|
|external_login_federation_provider_url|String|
|external_login_id|String|
|identity_type|String|
|principal_id|String|
|role|String|
|user_name|String|