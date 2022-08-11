
# Table: cloudflare_waf_rules
WAFRule represents a WAF rule.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|waf_cq_id|uuid|Unique CloudQuery ID of cloudflare_waf table (FK)|
|account_id|text|The Account ID of the resource.|
|zone_id|text|Zone identifier tag.|
|group|jsonb|The rule group to which the current WAF rule belongs.|
|id|text|The unique identifier of the WAF rule.|
|description|text|The public description of the WAF rule.|
|priority|text|The order in which the individual WAF rule is executed within its rule group.|
|package_id|text||
|mode|text|The action that the current WAF rule will perform when triggered. Applies to traditional (deny) WAF rules.|
|default_mode|text|The default action/mode of a rule.|
|allowed_modes|text[]|The list of possible actions of the WAF rule when it is triggered.|
