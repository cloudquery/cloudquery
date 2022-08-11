
# Table: aws_wafv2_web_acls
A Web ACL defines a collection of rules to use to inspect and control web requests
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|scope|text|Specifies whether this is for an Amazon CloudFront distribution or for a regional application.|
|resources_for_web_acl|text[]||
|tags|jsonb||
|arn|text|The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.  |
|default_action|jsonb|The action to perform if none of the Rules contained in the WebACL match.  |
|id|text|A unique identifier for the WebACL|
|name|text|The name of the Web ACL|
|visibility_config_cloud_watch_metrics_enabled|boolean|A boolean indicating whether the associated resource sends metrics to CloudWatch|
|visibility_config_metric_name|text|A name of the CloudWatch metric|
|visibility_config_sampled_requests_enabled|boolean|A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules|
|capacity|bigint|The web ACL capacity units (WCUs) currently being used by this web ACL|
|custom_response_bodies|jsonb|A map of custom response keys and content bodies|
|description|text|A description of the Web ACL that helps with identification.|
|label_namespace|text|The label namespace prefix for this web ACL|
|managed_by_firewall_manager|boolean|Indicates whether this web ACL is managed by AWS Firewall Manager|
|logging_configuration|text[]|The LoggingConfiguration for the specified web ACL.|
