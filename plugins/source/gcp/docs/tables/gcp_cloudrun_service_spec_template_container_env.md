
# Table: gcp_cloudrun_service_spec_template_container_env
EnvVar represents an environment variable present in a Container
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_spec_template_container_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_service_spec_template_containers table (FK)|
|name|text|Name of the environment variable|
|value|text|Variable references $(VAR_NAME) are expanded using the previous defined environment variables in the container and any route environment variables|
|value_from_config_map_key_ref_key|text|The key to select|
|value_from_config_map_key_ref_local_object_reference_name|text|Name of the referent|
|value_from_config_map_key_ref_name|text|The ConfigMap to select from|
|value_from_config_map_key_ref_optional|boolean|Specify whether the ConfigMap or its key must be defined|
|value_from_secret_key_ref_key|text|A Cloud Secret Manager secret version|
|value_from_secret_key_ref_local_object_reference_name|text|Name of the referent|
|value_from_secret_key_ref_name|text|The name of the secret in Cloud Secret Manager|
|value_from_secret_key_ref_optional|boolean|Specify whether the Secret or its key must be defined|
