# Table: aws_wafv2_web_acls

This table shows data for Wafv2 Web ACLs.

https://docs.aws.amazon.com/waf/latest/APIReference/API_WebACL.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|resources_for_web_acl|`list<item: utf8, nullable>`|
|arn (PK)|`utf8`|
|default_action|`json`|
|id|`utf8`|
|name|`utf8`|
|visibility_config|`json`|
|capacity|`int64`|
|captcha_config|`json`|
|challenge_config|`json`|
|custom_response_bodies|`json`|
|description|`utf8`|
|label_namespace|`utf8`|
|managed_by_firewall_manager|`bool`|
|post_process_firewall_manager_rule_groups|`json`|
|pre_process_firewall_manager_rule_groups|`json`|
|rules|`json`|
|token_domains|`list<item: utf8, nullable>`|
|logging_configuration|`json`|