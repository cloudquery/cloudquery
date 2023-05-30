# Table: aws_iam_virtual_mfa_devices

This table shows data for IAM Virtual MFA Devices.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html

The primary key for this table is **serial_number**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|serial_number (PK)|`utf8`|
|tags|`json`|
|base32_string_seed|`binary`|
|enable_date|`timestamp[us, tz=UTC]`|
|qr_code_png|`binary`|
|user|`json`|