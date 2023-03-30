# Table: fastly_account_users

This table shows data for Fastly Account Users.

https://developer.fastly.com/reference/api/account/user/

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|created_at|Timestamp|
|customer_id|String|
|deleted_at|Timestamp|
|email_hash|String|
|limit_services|Bool|
|locked|Bool|
|login|String|
|name|String|
|require_new_password|Bool|
|role|String|
|two_factor_auth_enabled|Bool|
|two_factor_setup_required|Bool|
|updated_at|Timestamp|