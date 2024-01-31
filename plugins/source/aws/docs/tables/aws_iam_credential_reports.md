# Table: aws_iam_credential_reports

This table shows data for IAM Credential Reports.

https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_getting-report.html#id_credentials_understanding_the_report_format

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**arn**, **user_creation_time**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|arn|`utf8`|
|user_creation_time|`timestamp[us, tz=UTC]`|
|password_last_changed|`timestamp[us, tz=UTC]`|
|password_next_rotation|`timestamp[us, tz=UTC]`|
|access_key_1_last_rotated|`timestamp[us, tz=UTC]`|
|access_key_2_last_rotated|`timestamp[us, tz=UTC]`|
|cert_1_last_rotated|`timestamp[us, tz=UTC]`|
|cert_2_last_rotated|`timestamp[us, tz=UTC]`|
|access_key_1_last_used_date|`timestamp[us, tz=UTC]`|
|access_key_2_last_used_date|`timestamp[us, tz=UTC]`|
|password_last_used|`timestamp[us, tz=UTC]`|
|password_enabled|`utf8`|
|user|`utf8`|
|password_status|`utf8`|
|mfa_active|`bool`|
|access_key1_active|`bool`|
|access_key2_active|`bool`|
|cert1_active|`bool`|
|cert2_active|`bool`|
|access_key1_last_used_region|`utf8`|
|access_key1_last_used_service|`utf8`|
|access_key2_last_used_region|`utf8`|
|access_key2_last_used_service|`utf8`|