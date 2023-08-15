# Table: aws_apigateway_domain_names

This table shows data for Amazon API Gateway Domain Names.

https://docs.aws.amazon.com/apigateway/latest/api/API_DomainName.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_apigateway_domain_names:
  - [aws_apigateway_domain_name_base_path_mappings](aws_apigateway_domain_name_base_path_mappings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|certificate_arn|`utf8`|
|certificate_name|`utf8`|
|certificate_upload_date|`timestamp[us, tz=UTC]`|
|distribution_domain_name|`utf8`|
|distribution_hosted_zone_id|`utf8`|
|domain_name|`utf8`|
|domain_name_status|`utf8`|
|domain_name_status_message|`utf8`|
|endpoint_configuration|`json`|
|mutual_tls_authentication|`json`|
|ownership_verification_certificate_arn|`utf8`|
|regional_certificate_arn|`utf8`|
|regional_certificate_name|`utf8`|
|regional_domain_name|`utf8`|
|regional_hosted_zone_id|`utf8`|
|security_policy|`utf8`|
|tags|`json`|