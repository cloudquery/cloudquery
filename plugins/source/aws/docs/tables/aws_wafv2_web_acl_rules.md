
# Table: aws_wafv2_web_acl_rules
A single rule, which you can use in a WebACL or RuleGroup to identify web requests that you want to allow, block, or count
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_cq_id|uuid|Unique CloudQuery ID of aws_wafv2_web_acls table (FK)|
|name|text|The name of the rule|
|priority|integer|If you define more than one Rule in a WebACL, AWS WAF evaluates each request against the Rules in order based on the value of Priority|
|statement|jsonb|The AWS WAF processing statement for the rule, for example ByteMatchStatement or SizeConstraintStatement.  |
|visibility_config_cloud_watch_metrics_enabled|boolean|A boolean indicating whether the associated resource sends metrics to CloudWatch|
|visibility_config_metric_name|text|A name of the CloudWatch metric|
|visibility_config_sampled_requests_enabled|boolean|A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules|
|action|jsonb|The action that AWS WAF should take on a web request when it matches the rule statement|
|override_action|jsonb|The override action to apply to the rules in a rule group|
|labels|text[]|Labels to apply to web requests that match the rule match statement|
