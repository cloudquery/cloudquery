
# Table: azure_front_door_load_balancing_settings
Load balancing settings for a backend pool associated with the Front Door instance
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|front_door_cq_id|uuid|Unique CloudQuery ID of azure_front_doors table (FK)|
|resource_state|text|Resource state|
|sample_size|integer|The number of samples to consider for load balancing decisions|
|successful_samples_required|integer|The number of samples within the sample period that must succeed|
|additional_latency_milliseconds|integer|The additional latency in milliseconds for probes to fall into the lowest latency bucket|
|name|text|Resource name|
|type|text|Resource type|
|id|text|Resource ID|
