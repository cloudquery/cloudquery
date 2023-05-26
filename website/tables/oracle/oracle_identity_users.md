# Table: oracle_identity_users

This table shows data for Oracle Identity Users.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|description|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|lifecycle_state|`utf8`|
|is_mfa_activated|`bool`|
|email|`utf8`|
|email_verified|`bool`|
|db_user_name|`utf8`|
|identity_provider_id|`utf8`|
|external_identifier|`utf8`|
|inactive_status|`int64`|
|freeform_tags|`json`|
|defined_tags|`json`|
|capabilities|`json`|
|last_successful_login_time|`timestamp[us, tz=UTC]`|
|previous_successful_login_time|`timestamp[us, tz=UTC]`|