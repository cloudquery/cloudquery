# Table: aws_apigateway_api_keys

This table shows data for Amazon API Gateway API Keys.

https://docs.aws.amazon.com/apigateway/latest/api/API_ApiKey.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|account_id|utf8|
|region|utf8|
|arn (PK)|utf8|
|created_date|timestamp[us, tz=UTC]|
|customer_id|utf8|
|description|utf8|
|enabled|bool|
|id|utf8|
|last_updated_date|timestamp[us, tz=UTC]|
|name|utf8|
|stage_keys|list<item: utf8, nullable>|
|tags|json|
|value|utf8|