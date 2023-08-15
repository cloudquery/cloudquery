# Table: aws_iam_signing_certificates

This table shows data for IAM Signing Certificates.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_SigningCertificate.html

The composite primary key for this table is (**account_id**, **user_arn**, **certificate_id**).

## Relations

This table depends on [aws_iam_users](aws_iam_users).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|user_arn (PK)|`utf8`|
|user_id|`utf8`|
|certificate_body|`utf8`|
|certificate_id (PK)|`utf8`|
|status|`utf8`|
|user_name|`utf8`|
|upload_date|`timestamp[us, tz=UTC]`|