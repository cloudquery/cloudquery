
# Table: cloudflare_waf_overrides
WAFOverride represents a WAF override.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The Account ID of the resource.|
|zone_id|text|The Zone ID of the resource.|
|id|text|The unique identifier of the WAF override.|
|description|text|An informative summary of the current URI-based WAF override.|
|url_s|text[]||
|priority|bigint|The relative priority of the current URI-based WAF override when multiple overrides match a single URL. A lower number indicates higher priority. Higher priority overrides may overwrite values set by lower priority overrides.|
|groups|jsonb|An object that allows you to enable or disable WAF rule groups for the current WAF override. Each key of this object must be the ID of a WAF rule group, and each value must be a valid WAF action (usually default or disable). When creating a new URI-based WAF override, you must provide a groups object or a rules object.|
|rewrite_action|jsonb|Specifies that, when a WAF rule matches, its configured action will be replaced by the action configured in this object.|
|rules|jsonb|The default action performed by the rules in the WAF package.|
|paused|boolean|When true, indicates that the WAF package is currently paused.|
