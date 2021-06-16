
# Table: aws_cloudfront_distribution_origins

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|distribution_id|uuid||
|domain_name|text||
|origin_id|text||
|connection_attempts|integer||
|connection_timeout|integer||
|custom_headers|jsonb||
|custom_origin_config_http_port|integer||
|custom_origin_config_https_port|integer||
|custom_origin_config_protocol_policy|text||
|custom_origin_config_keepalive_timeout|integer||
|custom_origin_config_read_timeout|integer||
|custom_origin_config_ssl_protocols|text[]||
|origin_path|text||
|origin_shield_enabled|boolean||
|origin_shield_region|text||
|s3_origin_config_origin_access_identity|text||
