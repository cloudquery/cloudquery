
# Table: aws_access_analyzer_analyzers
Contains information about the analyzer.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The ARN of the analyzer.|
|created_at|timestamp without time zone|A timestamp for the time at which the analyzer was created.|
|name|text|The name of the analyzer.|
|status|text|The status of the analyzer. An Active analyzer successfully monitors supported resources and generates new findings. The analyzer is Disabled when a user action, such as removing trusted access for AWS IAM Access Analyzer from AWS Organizations, causes the analyzer to stop generating new findings. The status is Creating when the analyzer creation is in progress and Failed when the analyzer creation has failed.|
|type|text|The type of analyzer, which corresponds to the zone of trust chosen for the analyzer.|
|last_resource_analyzed|text|The resource that was most recently analyzed by the analyzer.|
|last_resource_analyzed_at|timestamp without time zone|The time at which the most recently analyzed resource was analyzed.|
|status_reason_code|text|The reason code for the current status of the analyzer.|
|tags|jsonb|The tags added to the analyzer.|
