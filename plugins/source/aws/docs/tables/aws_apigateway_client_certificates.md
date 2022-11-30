# Table: aws_apigateway_client_certificates

https://docs.aws.amazon.com/apigateway/latest/api/API_ClientCertificate.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|client_certificate_id|String|
|created_date|Timestamp|
|description|String|
|expiration_date|Timestamp|
|pem_encoded_certificate|String|
|tags|JSON|