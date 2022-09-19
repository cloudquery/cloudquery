
# Table: azure_network_route_filter_rules
Route Filter Rule Resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|route_filter_cq_id|uuid|Unique CloudQuery ID of azure_network_route_filters table (FK)|
|id|text|Resource ID.|
|access|text|The access type of the rule.|
|communities|text[]|The collection for bgp community values to filter on.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|location|text|Resource location.|
|name|text|Resource name.|
|provisioning_state|text|The provisioning state of the route filter rule resource.|
|route_filter_rule_type|text|The rule type of the rule.|
