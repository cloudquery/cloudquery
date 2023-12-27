# Table: aws_quicksight_users

This table shows data for QuickSight Users.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_User.html

The composite primary key for this table is (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|tags|`json`|
|active|`bool`|
|arn (PK)|`utf8`|
|custom_permissions_name|`utf8`|
|email|`utf8`|
|external_login_federation_provider_type|`utf8`|
|external_login_federation_provider_url|`utf8`|
|external_login_id|`utf8`|
|identity_type|`utf8`|
|principal_id|`utf8`|
|role|`utf8`|
|user_name|`utf8`|