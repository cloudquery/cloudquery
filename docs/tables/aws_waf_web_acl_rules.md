
# Table: aws_waf_web_acl_rules
This is AWS WAF Classic documentation
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_cq_id|uuid|Unique CloudQuery ID of aws_waf_web_acls table (FK)|
|priority|integer|Specifies the order in which the Rules in a WebACL are evaluated|
|rule_id|text|The RuleId for a Rule|
|action_type|text|Specifies how you want AWS WAF to respond to requests that match the settings in a Rule|
|excluded_rules|text[]|An array of rules to exclude from a rule group|
|override_action_type|text|COUNT overrides the action specified by the individual rule within a RuleGroup . If set to NONE, the rule's action will take place.  |
|type|text|The rule type, either REGULAR, as defined by Rule, RATE_BASED, as defined by RateBasedRule, or GROUP, as defined by RuleGroup|
