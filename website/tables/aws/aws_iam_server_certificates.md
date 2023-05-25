# Table: aws_iam_server_certificates

This table shows data for IAM Server Certificates.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServerCertificateMetadata.html

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id (PK)|utf8|
|id (PK)|utf8|
|arn|utf8|
|path|utf8|
|server_certificate_id|utf8|
|server_certificate_name|utf8|
|expiration|timestamp[us, tz=UTC]|
|upload_date|timestamp[us, tz=UTC]|