# Table: aws_apigateway_client_certificates

This table shows data for Amazon API Gateway Client Certificates.

https://docs.aws.amazon.com/apigateway/latest/api/API_ClientCertificate.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|client_certificate_id|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|description|`utf8`|
|expiration_date|`timestamp[us, tz=UTC]`|
|pem_encoded_certificate|`utf8`|
|tags|`json`|