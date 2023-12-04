# Table: aws_identitystore_users

This table shows data for Identity Store Users.

https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html

The composite primary key for this table is (**request_account_id**, **request_region**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id (PK)|`utf8`|
|request_region (PK)|`utf8`|
|arn (PK)|`utf8`|
|identity_store_id|`utf8`|
|user_id|`utf8`|
|addresses|`json`|
|display_name|`utf8`|
|emails|`json`|
|external_ids|`json`|
|locale|`utf8`|
|name|`json`|
|nick_name|`utf8`|
|phone_numbers|`json`|
|preferred_language|`utf8`|
|profile_url|`utf8`|
|timezone|`utf8`|
|title|`utf8`|
|user_name|`utf8`|
|user_type|`utf8`|