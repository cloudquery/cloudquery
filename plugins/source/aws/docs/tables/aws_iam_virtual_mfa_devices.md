# Table: aws_iam_virtual_mfa_devices

https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html

The primary key for this table is **serial_number**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|serial_number (PK)|String|
|tags|JSON|
|user_tags|JSON|
|base32_string_seed|IntArray|
|enable_date|Timestamp|
|qr_code_png|IntArray|
|user|JSON|