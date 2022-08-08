
# Table: azure_front_door_health_probe_settings
Health probe settings for a backend pool associated with this Front Door instance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_cq_id|uuid|Unique CloudQuery ID of azure_front_doors table (FK)|
|resource_state|text|Resource state|
|path|text|The path to use for the health probe|
|protocol|text|Protocol scheme to use for the health probe|
|interval_in_seconds|integer|The number of seconds between health probes|
|health_probe_method|text|Which HTTP method is used to perform the health probe|
|enabled_state|text|Whether the health probe is enabled|
|name|text|Resource name|
|type|text|Resource type|
|id|text|Resource ID|
