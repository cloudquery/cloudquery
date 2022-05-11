
# Table: gcp_serviceusage_service_quota_limits
`QuotaLimit` defines a specific limit that applies over a specified duration for a limit type
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_serviceusage_services table (FK)|
|default_limit|integer|Default number of tokens that can be consumed during the specified duration|
|description|text|Optional|
|display_name|text|User-visible display name for this limit|
|duration|text|Duration of this limit in textual notation|
|free_tier|bigint|Free tier value displayed in the Developers Console for this limit|
|max_limit|bigint|Maximum number of tokens that can be consumed during the specified duration|
|metric|text|The name of the metric this quota limit applies to|
|name|text|Name of the quota limit|
|unit|text|Specify the unit of the quota limit|
|values|jsonb|Tiered limit values|
