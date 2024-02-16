# Table: aws_iam_mfa_devices

This table shows data for IAM MFA Devices.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_MFADevice.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **serial_number**.
## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|serial_number|`utf8`|
|enable_date|`timestamp[us, tz=UTC]`|
|user_name|`utf8`|