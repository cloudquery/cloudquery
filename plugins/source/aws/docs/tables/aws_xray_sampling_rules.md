
# Table: aws_xray_sampling_rules
A SamplingRule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb|A list of Tags that specify information about the sampling rule.|
|created_at|timestamp without time zone|When the rule was created.|
|modified_at|timestamp without time zone|When the rule was last modified.|
|fixed_rate|float|The percentage of matching requests to instrument, after the reservoir is exhausted.|
|http_method|text|Matches the HTTP method of a request.|
|host|text|Matches the hostname from a request URL.|
|priority|integer|The priority of the sampling rule.|
|reservoir_size|integer|A fixed number of matching requests to instrument per second, prior to applying the fixed rate|
|resource_arn|text|Matches the ARN of the Amazon Web Services resource on which the service runs.|
|service_name|text|Matches the name that the service uses to identify itself in segments.|
|service_type|text|Matches the origin that the service uses to identify its type in segments.|
|url_path|text|Matches the path from a request URL.|
|version|integer|The version of the sampling rule format (1).|
|attributes|jsonb|Matches attributes derived from the request.|
|arn|text|The ARN of the sampling rule|
|rule_name|text|The name of the sampling rule|
