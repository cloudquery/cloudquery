# Table: aws_iam_virtual_mfa_devices

This table shows data for IAM Virtual MFA Devices.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **serial_number**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|serial_number|`utf8`|
|tags|`json`|
|base32_string_seed|`binary`|
|enable_date|`timestamp[us, tz=UTC]`|
|qr_code_png|`binary`|
|user|`json`|