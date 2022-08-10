
# Table: aws_wafv2_web_acl_logging_configuration
The LoggingConfiguration for the specified web ACL.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_cq_id|uuid|Unique CloudQuery ID of aws_wafv2_web_acls table (FK)|
|log_destination_configs|text[]|The Amazon Resource Names (ARNs) of the logging destinations that you want to associate with the web ACL.|
|resource_arn|text|The Amazon Resource Name (ARN) of the web ACL that you want to associate with LogDestinationConfigs.|
|logging_filter|jsonb|Filtering that specifies which web requests are kept in the logs and which are dropped. You can filter on the rule action and on the web request labels that were applied by matching rules during web ACL evaluation.|
|managed_by_firewall_manager|boolean|Indicates whether the logging configuration was created by Firewall Manager, as part of an WAF policy configuration. If true, only Firewall Manager can modify or delete the configuration.|
|redacted_fields|jsonb|The parts of the request that you want redacted from the logs. For example, if you redact the cookie field, the cookie field in the firehose will be xxx. You can specify only the following fields for redaction: UriPath, QueryString, SingleHeader, Method, and JsonBody.|
