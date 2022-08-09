
# Table: cloudflare_waf
WAFPackage represents a WAF package configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The Account ID of the resource.|
|id|text|The unique identifier of a WAF package.|
|name|text|The name of the WAF package.|
|description|text|A summary of the purpose/function of the WAF package.|
|zone_id|text|Zone identifier tag.|
|detection_mode|text|When a WAF package uses anomaly detection, each rule is given a score when triggered. If the total score of all triggered rules exceeds the sensitivity defined on the WAF package, the action defined on the package will be taken.|
|sensitivity|text|The sensitivity of the WAF package|
|action_mode|text|The default action performed by the rules in the WAF package.|
