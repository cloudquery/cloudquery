
# Table: azure_front_door_backend_pools
Backend pools available to routing rules
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_cq_id|uuid|Unique CloudQuery ID of azure_front_doors table (FK)|
|resource_state|text|Resource state|
|load_balancing_settings_id|text|Load balancing settings ID for the backend pool|
|health_probe_settings_id|text|L7 health probe settings ID for the backend pool|
|name|text|Resource name|
|type|text|Resource type|
|id|text|Resource ID|
