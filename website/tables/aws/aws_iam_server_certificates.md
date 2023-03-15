# Table: aws_iam_server_certificates

This table shows data for IAM Server Certificates.

https://docs.aws.amazon.com/IAM/latest/APIReference/API_ServerCertificateMetadata.html

The composite primary key for this table is (**account_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id (PK)|String|
|id (PK)|String|
|arn|String|
|path|String|
|server_certificate_id|String|
|server_certificate_name|String|
|expiration|Timestamp|
|upload_date|Timestamp|