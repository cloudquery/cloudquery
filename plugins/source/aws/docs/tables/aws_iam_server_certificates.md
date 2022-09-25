# Table: aws_iam_server_certificates


The composite primary key for this table is (**account_id**, **id**).


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id (PK)|String|
|id (PK)|String|
|arn|String|
|path|String|
|server_certificate_name|String|
|expiration|Timestamp|
|upload_date|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|