# Table: aws_identitystore_users

This table shows data for Identity Store Users.

https://docs.aws.amazon.com/singlesignon/latest/IdentityStoreAPIReference/API_User.html

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
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