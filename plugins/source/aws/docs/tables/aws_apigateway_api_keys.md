# Table: aws_apigateway_api_keys

This table shows data for Amazon API Gateway API Keys.

https://docs.aws.amazon.com/apigateway/latest/api/API_ApiKey.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|created_date|`timestamp[us, tz=UTC]`|
|customer_id|`utf8`|
|description|`utf8`|
|enabled|`bool`|
|id|`utf8`|
|last_updated_date|`timestamp[us, tz=UTC]`|
|name|`utf8`|
|stage_keys|`list<item: utf8, nullable>`|
|tags|`json`|
|value|`utf8`|