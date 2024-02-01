# Table: aws_servicequotas_quotas

This table shows data for Servicequotas Quotas.

https://docs.aws.amazon.com/servicequotas/2019-06-24/apireference/API_ServiceQuota.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

This table depends on [aws_servicequotas_services](aws_servicequotas_services.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|adjustable|`bool`|
|error_reason|`json`|
|global_quota|`bool`|
|period|`json`|
|quota_applied_at_level|`utf8`|
|quota_arn|`utf8`|
|quota_code|`utf8`|
|quota_context|`json`|
|quota_name|`utf8`|
|service_code|`utf8`|
|service_name|`utf8`|
|unit|`utf8`|
|usage_metric|`json`|
|value|`float64`|