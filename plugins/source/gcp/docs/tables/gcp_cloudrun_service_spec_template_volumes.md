
# Table: gcp_cloudrun_service_spec_template_volumes
Volume represents a named volume in a container
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_services table (FK)|
|config_map_default_mode|bigint|Integer representation of mode bits to use on created files by default|
|config_map_name|text|Name of the config|
|config_map_optional|boolean|Specify whether the Secret or its keys must be defined|
|name|text|Volume's name|
|secret_default_mode|bigint|Integer representation of mode bits to use on created files by default|
|secret_optional|boolean|Specify whether the Secret or its keys must be defined|
|secret_name|text|The name of the secret in Cloud Secret Manager|
