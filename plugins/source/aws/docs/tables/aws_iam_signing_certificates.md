# Table: aws_iam_signing_certificates

This table shows data for IAM Signing Certificates.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_SigningCertificate.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **user_arn**, **certificate_id**).
## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|user_arn|`utf8`|
|user_id|`utf8`|
|certificate_body|`utf8`|
|certificate_id|`utf8`|
|status|`utf8`|
|user_name|`utf8`|
|upload_date|`timestamp[us, tz=UTC]`|