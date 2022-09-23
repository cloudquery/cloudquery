# Table: aws_apigateway_client_certificates


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|client_certificate_id|String|
|created_date|Timestamp|
|description|String|
|expiration_date|Timestamp|
|pem_encoded_certificate|String|
|tags|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|