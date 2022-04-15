
# Table: aws_wafregional_web_acl_rules
The action for each Rule in a WebACL
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|web_acl_cq_id|uuid|Unique CloudQuery ID of aws_wafregional_web_acls table (FK)|
|priority|integer|Specifies the order in which the Rules in a WebACL are evaluated|
|rule_id|text|The RuleId for a Rule|
|action|text|Specifies how you want AWS WAF to respond to requests that match the settings in a Rule|
|excluded_rules|text[]|An array of rules to exclude from a rule group|
|override_action|text|Describes an override action for the rule.|
|type|text|The rule type, either REGULAR, as defined by Rule, RATE_BASED, as defined by RateBasedRule, or GROUP, as defined by RuleGroup|
