# Table: aws_apprunner_custom_domains

This table shows data for AWS App Runner Custom Domains.

https://docs.aws.amazon.com/apprunner/latest/api/API_CustomDomain.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**service_arn**, **domain_name**).
## Relations

This table depends on [aws_apprunner_services](aws_apprunner_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|service_arn|`utf8`|
|enable_www_subdomain|`bool`|
|domain_name|`utf8`|
|status|`utf8`|
|certificate_validation_records|`json`|