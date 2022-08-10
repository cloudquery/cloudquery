
# Table: aws_waf_web_acl_logging_configuration
The LoggingConfiguration for the specified web ACL.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_cq_id|uuid|Unique CloudQuery ID of aws_waf_web_acls table (FK)|
|log_destination_configs|text[]|An array of Amazon Kinesis Data Firehose ARNs.|
|resource_arn|text|The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.|
|redacted_fields|jsonb|The parts of the request that you want redacted from the logs. For example, if you redact the cookie field, the cookie field in the firehose will be xxx.|
