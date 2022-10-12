# Table: aws_rds_certificates



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
|certificate_identifier|String|
|certificate_type|String|
|customer_override|Bool|
|customer_override_valid_till|Timestamp|
|thumbprint|String|
|valid_from|Timestamp|
|valid_till|Timestamp|