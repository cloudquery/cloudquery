
# Table: cloudflare_waf_rule_groups
WAFGroup represents a WAF rule group.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|waf_cq_id|uuid|Unique CloudQuery ID of cloudflare_waf table (FK)|
|account_id|text|The Account ID of the resource.|
|zone_id|text|Zone identifier tag.|
|id|text|The unique identifier of the rule group.|
|name|text|The name of the rule group.|
|description|text|An informative summary of what the rule group does.|
|rules_count|bigint|The number of rules in the current rule group.|
|modified_rules_count|bigint|The number of rules within the group that have been modified from their default configuration.|
|package_id|text|The unique identifier of a WAF package.|
|mode|text|The state of the rules contained in the rule group. When on, the rules in the group are configurable/usable.|
|allowed_modes|text[]|The available states for the rule group.|
