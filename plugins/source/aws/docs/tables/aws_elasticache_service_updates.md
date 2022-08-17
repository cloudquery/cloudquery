
# Table: aws_elasticache_service_updates
An update that you can apply to your Redis clusters.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|auto_update_after_recommended_apply_by_date|boolean|Indicates whether the service update will be automatically applied once the recommended apply-by date has expired.|
|engine|text|The Elasticache engine to which the update applies|
|engine_version|text|The Elasticache engine version to which the update applies|
|estimated_update_time|text|The estimated length of time the service update will take|
|description|text|Provides details of the service update|
|end_date|timestamp without time zone|The date after which the service update is no longer available|
|name|text|The unique ID of the service update|
|recommended_apply_by_date|timestamp without time zone|The recommendend date to apply the service update in order to ensure compliance. For information on compliance, see Self-Service Security Updates for Compliance (https://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/elasticache-compliance.html#elasticache-compliance-self-service).|
|release_date|timestamp without time zone|The date when the service update is initially available|
|severity|text|The severity of the service update|
|status|text|The status of the service update|
|type|text|Reflects the nature of the service update|
