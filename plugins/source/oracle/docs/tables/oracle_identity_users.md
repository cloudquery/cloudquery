# Table: oracle_identity_users

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|compartment_id|String|
|name|String|
|description|String|
|time_created|Timestamp|
|lifecycle_state|String|
|is_mfa_activated|Bool|
|email|String|
|email_verified|Bool|
|db_user_name|String|
|identity_provider_id|String|
|external_identifier|String|
|inactive_status|Int|
|freeform_tags|JSON|
|defined_tags|JSON|
|capabilities|JSON|
|last_successful_login_time|Timestamp|
|previous_successful_login_time|Timestamp|