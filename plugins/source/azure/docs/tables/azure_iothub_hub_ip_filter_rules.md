
# Table: azure_iothub_hub_ip_filter_rules
IPFilterRule the IP filter rules for the IoT hub.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|hub_cq_id|uuid|Unique CloudQuery ID of azure_iothub_hubs table (FK)|
|filter_name|text|The name of the IP filter rule.|
|action|text|The desired action for requests captured by this rule|
|ip_mask|text|A string that contains the IP address range in CIDR notation for the rule.|
