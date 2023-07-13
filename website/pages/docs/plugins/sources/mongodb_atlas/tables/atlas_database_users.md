# Table: atlas_database_users

This table shows data for Atlas Database Users.

The composite primary key for this table is (**database_name**, **group_id**, **username**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|aws_iam_type|`utf8`|
|database_name (PK)|`utf8`|
|delete_after_date|`timestamp[us, tz=UTC]`|
|group_id (PK)|`utf8`|
|labels|`json`|
|ldap_auth_type|`utf8`|
|links|`json`|
|oidc_auth_type|`utf8`|
|password|`utf8`|
|roles|`json`|
|scopes|`json`|
|username (PK)|`utf8`|
|x509_type|`utf8`|