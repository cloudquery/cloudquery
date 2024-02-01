# Table: aws_quicksight_users

This table shows data for QuickSight Users.

https://docs.aws.amazon.com/quicksight/latest/APIReference/API_User.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|active|`bool`|
|arn|`utf8`|
|custom_permissions_name|`utf8`|
|email|`utf8`|
|external_login_federation_provider_type|`utf8`|
|external_login_federation_provider_url|`utf8`|
|external_login_id|`utf8`|
|identity_type|`utf8`|
|principal_id|`utf8`|
|role|`utf8`|
|user_name|`utf8`|