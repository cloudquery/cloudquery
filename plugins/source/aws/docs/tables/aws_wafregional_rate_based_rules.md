
# Table: aws_wafregional_rate_based_rules
A combination of identifiers for web requests that you want to allow, block, or count, including rate limit.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|ARN of the rate based rule.|
|tags|jsonb|Rule tags.|
|rate_key|text|The field that AWS WAF uses to determine if requests are likely arriving from single source and thus subject to rate monitoring|
|rate_limit|bigint|The maximum number of requests, which have an identical value in the field specified by the RateKey, allowed in a five-minute period|
|id|text|A unique identifier for a RateBasedRule|
|metric_name|text|A friendly name or description for the metrics for a RateBasedRule|
|name|text|A friendly name or description for a RateBasedRule|
