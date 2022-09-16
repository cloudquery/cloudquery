
# Table: aws_apigateway_usage_plans
Represents a usage plan used to specify who can assess associated API stages
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource|
|region|text|The AWS Region of the resource|
|arn|text|The Amazon Resource Name (ARN) for the resource|
|description|text|The description of a usage plan|
|id|text|The identifier of a UsagePlan resource|
|name|text|The name of a usage plan|
|product_code|text|The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace|
|quota_limit|bigint|The target maximum number of requests that can be made in a given time period|
|quota_offset|bigint|The number of requests subtracted from the given limit in the initial time period|
|quota_period|text|The time period in which the limit applies|
|tags|jsonb|The collection of tags|
|throttle_burst_limit|bigint|The API target request burst rate limit|
|throttle_rate_limit|float|The API target request rate limit|
