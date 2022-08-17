
# Table: azure_cdn_profile_endpoint_origin_groups
DeepCreatedOriginGroup the origin group for CDN content which is added when creating a CDN endpoint Traffic is sent to the origins within the origin group based on origin health
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|profile_endpoint_cq_id|uuid|Unique CloudQuery ID of azure_cdn_profile_endpoints table (FK)|
|name|text|Origin group name which must be unique within the endpoint|
|health_probe_settings_probe_path|text|The path relative to the origin that is used to determine the health of the origin|
|health_probe_settings_probe_request_type|text|The type of health probe request that is made|
|health_probe_settings_probe_protocol|text|Protocol to use for health probe|
|health_probe_settings_probe_interval_in_seconds|bigint|The number of seconds between health probesDefault is 240sec|
|origins|text[]|The source of the content being delivered via CDN within given origin group|
|traffic_restoration_time_to_healed_or_new_endpoints_in_minutes|bigint|Time in minutes to shift the traffic to the endpoint gradually when an unhealthy endpoint comes healthy or a new endpoint is added|
|response_based_origin_error_detection_settings|jsonb|The JSON object that contains the properties to determine origin health using real requests/responsesThis property is currently not supported|
