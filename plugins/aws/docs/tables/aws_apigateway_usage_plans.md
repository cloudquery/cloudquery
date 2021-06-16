
# Table: aws_apigateway_usage_plans
Represents a usage plan than can specify who can assess associated API stages with specified request limits and quotas.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|description|text|The description of a usage plan.|
|resource_id|text|The identifier of a UsagePlan resource.|
|name|text|The name of a usage plan.|
|product_code|text|The AWS Markeplace product identifier to associate with the usage plan as a SaaS product on AWS Marketplace.|
|quota_limit|integer|The maximum number of requests that can be made in a given time period.|
|quota_offset|integer|The day that a time period starts. For example, with a time period of WEEK, an offset of 0 starts on Sunday, and an offset of 1 starts on Monday.|
|quota_period|text|The time period in which the limit applies. Valid values are "DAY", "WEEK" or "MONTH".|
|tags|jsonb|The collection of tags. Each tag element is associated with a given resource.|
|throttle_burst_limit|integer|The API request burst limit, the maximum rate limit over a time ranging from one to a few seconds, depending upon whether the underlying token bucket is at its full capacity.|
|throttle_rate_limit|float|The API request steady-state rate limit.|
