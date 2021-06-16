
# Table: aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups
A rule group that's defined for an AWS Firewall Manager WAF policy. 
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_id|uuid|Unique ID of aws_wafv2_web_acls table (FK)|
|statement|jsonb|The processing guidance for an AWS Firewall Manager rule|
|name|text|The name of the rule group|
|override_action|jsonb|The override action to apply to the rules in a rule group|
|priority|integer|If you define more than one rule group in the first or last Firewall Manager rule groups, AWS WAF evaluates each request against the rule groups in order, starting from the lowest priority setting|
|visibility_config_cloud_watch_metrics_enabled|boolean|A boolean indicating whether the associated resource sends metrics to CloudWatch|
|visibility_config_metric_name|text|A name of the CloudWatch metric|
|visibility_config_sampled_requests_enabled|boolean|A boolean indicating whether AWS WAF should store a sampling of the web requests that match the rules|
