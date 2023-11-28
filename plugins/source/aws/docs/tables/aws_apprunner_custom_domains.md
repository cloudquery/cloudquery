# Table: aws_apprunner_custom_domains

This table shows data for AWS App Runner Custom Domains.

https://docs.aws.amazon.com/apprunner/latest/api/API_CustomDomain.html

The composite primary key for this table is (**service_arn**, **domain_name**).

## Relations

This table depends on [aws_apprunner_services](aws_apprunner_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|service_arn (PK)|`utf8`|
|enable_www_subdomain|`bool`|
|domain_name (PK)|`utf8`|
|status|`utf8`|
|certificate_validation_records|`json`|