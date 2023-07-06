# Table: aws_servicequotas_quotas

This table shows data for Servicequotas Quotas.

https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceQuota.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_servicequotas_services](aws_servicequotas_services).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|adjustable|`bool`|
|error_reason|`json`|
|global_quota|`bool`|
|period|`json`|
|quota_arn|`utf8`|
|quota_code|`utf8`|
|quota_name|`utf8`|
|service_code|`utf8`|
|service_name|`utf8`|
|unit|`utf8`|
|usage_metric|`json`|
|value|`float64`|