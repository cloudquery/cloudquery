# Table: aws_iam_signing_certificates

https://docs.aws.amazon.com/IAM/latest/APIReference/API_SigningCertificate.html

The composite primary key for this table is (**user_arn**, **certificate_id**).

## Relations

This table depends on [aws_iam_users](aws_iam_users.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|user_arn (PK)|String|
|user_id|String|
|certificate_body|String|
|certificate_id (PK)|String|
|status|String|
|user_name|String|
|upload_date|Timestamp|