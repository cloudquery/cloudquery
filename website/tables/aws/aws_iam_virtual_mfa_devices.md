# Table: aws_iam_virtual_mfa_devices

This table shows data for IAM Virtual MFA Devices.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_VirtualMFADevice.html

The primary key for this table is **serial_number**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|serial_number (PK)|`utf8`|
|tags|`json`|
|base32_string_seed|`binary`|
|enable_date|`timestamp[us, tz=UTC]`|
|qr_code_png|`binary`|
|user|`json`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### Ensure hardware MFA is enabled for the "root" account (Scored)

```sql
SELECT
  'Ensure hardware MFA is enabled for the "root" account (Scored)' AS title,
  split_part(cr.arn, ':', 5) AS account_id,
  cr.arn AS resource_id,
  CASE
  WHEN mfa.serial_number IS NULL OR cr.mfa_active = false THEN 'fail'
  WHEN mfa.serial_number IS NOT NULL AND cr.mfa_active = true THEN 'pass'
  END
    AS status
FROM
  aws_iam_credential_reports AS cr
  LEFT JOIN aws_iam_virtual_mfa_devices AS mfa ON mfa.user->>'Arn' = cr.arn
WHERE
  cr.user = '<root_account>'
GROUP BY
  mfa.serial_number, cr.mfa_active, cr.arn;
```


