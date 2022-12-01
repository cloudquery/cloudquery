# Table: aws_iot_ca_certificates

https://docs.aws.amazon.com/iot/latest/apireference/API_CACertificateDescription.html

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
|certificates|StringArray|
|arn (PK)|String|
|auto_registration_status|String|
|certificate_id|String|
|certificate_mode|String|
|certificate_pem|String|
|creation_date|Timestamp|
|customer_version|Int|
|generation_id|String|
|last_modified_date|Timestamp|
|owned_by|String|
|status|String|
|validity|JSON|