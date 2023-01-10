# Table: aws_apprunner_custom_domains

https://docs.aws.amazon.com/apprunner/latest/api/API_CustomDomain.html

The primary key for this table is **_cq_id**.

## Relations

This table depends on [aws_apprunner_services](aws_apprunner_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|enable_www_subdomain|Bool|
|domain_name|String|
|status|String|
|certificate_validation_records|JSON|