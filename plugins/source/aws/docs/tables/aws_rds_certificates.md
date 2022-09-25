# Table: aws_rds_certificates


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|arn (PK)|String|
|certificate_identifier|String|
|certificate_type|String|
|customer_override|Bool|
|customer_override_valid_till|Timestamp|
|thumbprint|String|
|valid_from|Timestamp|
|valid_till|Timestamp|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|