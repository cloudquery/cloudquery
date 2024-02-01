# Table: aws_iam_server_certificates

This table shows data for IAM Server Certificates.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServerCertificateMetadata.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|id|`utf8`|
|arn|`utf8`|
|path|`utf8`|
|server_certificate_id|`utf8`|
|server_certificate_name|`utf8`|
|expiration|`timestamp[us, tz=UTC]`|
|upload_date|`timestamp[us, tz=UTC]`|