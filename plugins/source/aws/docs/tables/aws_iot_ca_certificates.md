# Table: aws_iot_ca_certificates


The primary key for this table is **arn**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|certificates|StringArray|
|arn (PK)|String|
|auto_registration_status|String|
|certificate_id|String|
|certificate_pem|String|
|creation_date|Timestamp|
|customer_version|Int|
|generation_id|String|
|last_modified_date|Timestamp|
|owned_by|String|
|status|String|
|validity|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|