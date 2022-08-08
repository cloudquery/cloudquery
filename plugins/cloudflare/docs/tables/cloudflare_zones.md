
# Table: cloudflare_zones
Zone describes a Cloudflare zone.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The Account ID of the resource.|
|host_name|text|Zone host name.|
|host_website|text|Zone host website.|
|id|text|The unique universal identifier for a Cloudflare zone.|
|name|text|Cloudflare zone name.|
|dev_mode|bigint|DevMode contains the time in seconds until development expires (if positive) or since it expired (if negative)|
|original_ns|text[]|Creation timestamp of the account.|
|original_registrar|text|Cloudflare zone original name servers.|
|original_dns_host|text|Cloudflare zone original registrar.|
|created_on|timestamp without time zone|Cloudflare zone original dns host.|
|modified_on|timestamp without time zone|Zone created on date and time.|
|name_servers|text[]|Zone last modified date and time.|
|owner_id|text|Zone owner id.|
|owner_email|text|Zone owner email address.|
|owner_name|text|Zone owner name.|
|owner_type|text|Zone owner type.|
|permissions|text[]|Zone permissions.|
|plan_id|text|The unique universal identifier for a Cloudflare zone plan.|
|plan_name|text|Cloudflare zone plan name.|
|plan_price|bigint|Cloudflare zone plan price.|
|plan_currency|text|Cloudflare zone plan currency.|
|plan_frequency|text|Cloudflare zone plan frequency.|
|plan_legacy_id|text|True if zone plan is subscribed.|
|plan_is_subscribed|boolean|True if zone plan can subscribe.|
|plan_can_subscribe|boolean|Cloudflare zone plan legacy id.|
|plan_legacy_discount|boolean|True if zone plan has legacy discount.|
|plan_externally_managed|boolean|True if zone plan is externally managed.|
|plan_pending_id|text|The unique universal identifier for a Cloudflare zone plan.|
|plan_pending_name|text|Cloudflare zone plan name.|
|plan_pending_price|bigint|Cloudflare zone plan price.|
|plan_pending_currency|text|Cloudflare zone plan currency.|
|plan_pending_frequency|text|Cloudflare zone plan frequency.|
|plan_pending_legacy_id|text|True if zone plan is subscribed.|
|plan_pending_is_subscribed|boolean|True if zone plan can subscribe.|
|plan_pending_can_subscribe|boolean|Cloudflare zone plan legacy id.|
|plan_pending_legacy_discount|boolean|True if zone plan has legacy discount.|
|plan_pending_externally_managed|boolean|True if zone plan is externally managed.|
|status|text|Zone status.|
|paused|boolean|True if zone is paused.|
|type|text|Zone type.|
|vanity_ns|text[]|Zone vanity name servers.|
|betas|text[]|Zone betas.|
|deactivation_reason|text|Zone deactivation reason.|
|page_rule_quota|bigint|custom_certificate_quota is broken - sometimes it's a string, sometimes a number! CustCertQuota     int    `json:"custom_certificate_quota"`|
|wildcard_proxiable|boolean||
|phishing_detected|boolean||
|verification_key|text|Zone verification key.|
