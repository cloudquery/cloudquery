# Table: aws_eventbridge_api_destinations

This table shows data for Amazon EventBridge API Destinations.

https://docs.aws.amazon.com/eventbridge/latest/APIReference/API_ApiDestination.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|api_destination_arn|`utf8`|
|api_destination_state|`utf8`|
|connection_arn|`utf8`|
|creation_time|`timestamp[us, tz=UTC]`|
|http_method|`utf8`|
|invocation_endpoint|`utf8`|
|invocation_rate_limit_per_second|`int64`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|name|`utf8`|