# Table: aws_iam_credential_reports



The composite primary key for this table is (**arn**, **user_creation_time**).



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|arn (PK)|String|
|user_creation_time (PK)|Timestamp|
|password_last_changed|Timestamp|
|password_next_rotation|Timestamp|
|access_key_1_last_rotated|Timestamp|
|access_key_2_last_rotated|Timestamp|
|cert_1_last_rotated|Timestamp|
|cert_2_last_rotated|Timestamp|
|access_key_1_last_used_date|Timestamp|
|access_key_2_last_used_date|Timestamp|
|password_last_used|Timestamp|
|password_enabled|String|
|user|String|
|password_status|String|
|mfa_active|Bool|
|access_key1_active|Bool|
|access_key2_active|Bool|
|cert1_active|Bool|
|cert2_active|Bool|
|access_key1_last_used_region|String|
|access_key1_last_used_service|String|
|access_key2_last_used_region|String|
|access_key2_last_used_service|String|