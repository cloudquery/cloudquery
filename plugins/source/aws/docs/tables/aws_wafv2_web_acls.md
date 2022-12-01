# Table: aws_wafv2_web_acls



The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|tags|JSON|
|resources_for_web_acl|StringArray|
|arn (PK)|String|
|default_action|JSON|
|id|String|
|name|String|
|visibility_config|JSON|
|capacity|Int|
|captcha_config|JSON|
|challenge_config|JSON|
|custom_response_bodies|JSON|
|description|String|
|label_namespace|String|
|managed_by_firewall_manager|Bool|
|post_process_firewall_manager_rule_groups|JSON|
|pre_process_firewall_manager_rule_groups|JSON|
|rules|JSON|
|token_domains|StringArray|
|logging_configuration|JSON|