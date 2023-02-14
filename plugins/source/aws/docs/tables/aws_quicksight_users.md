# Table: aws_quicksight_users

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_User.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|region (PK)|String|
|tags|JSON|
|active|Bool|
|arn (PK)|String|
|custom_permissions_name|String|
|email|String|
|external_login_federation_provider_type|String|
|external_login_federation_provider_url|String|
|external_login_id|String|
|identity_type|String|
|principal_id|String|
|role|String|
|user_name|String|