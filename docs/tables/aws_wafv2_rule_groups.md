
# Table: aws_wafv2_rule_groups
A rule group defines a collection of rules to inspect and control web requests that you can use in a WebACL
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|tags|jsonb||
|policy|jsonb||
|arn|text|The Amazon Resource Name (ARN) of the entity.  |
|capacity|bigint|The web ACL capacity units (WCUs) required for this rule group|
|resource_id|text|A unique identifier for the rule group|
|name|text|The name of the rule group|
|visibility_config_cloud_watch_metrics_enabled|boolean|A boolean indicating whether the associated resource sends metrics to CloudWatch|
|visibility_config_metric_name|text|A name of the CloudWatch metric|
|visibility_config_sampled_requests_enabled|boolean|A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules|
|custom_response_bodies|jsonb|A map of custom response keys and content bodies|
|description|text|A description of the rule group that helps with identification.|
|label_namespace|text|The label namespace prefix for this rule group|
|rules|jsonb|The Rule statements used to identify the web requests that you want to allow, block, or count|
