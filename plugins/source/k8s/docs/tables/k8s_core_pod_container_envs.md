
# Table: k8s_core_pod_container_envs
EnvVar represents an environment variable present in a Container.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|pod_container_cq_id|uuid|Unique CloudQuery ID of k8s_core_pod_init_containers table (FK)|
|name|text|Name of the environment variable|
|value|text|Variable references $(VAR_NAME) are expanded using the previously defined environment variables in the container and any service environment variables|
|value_from_field_ref_api_version|text|Version of the schema the FieldPath is written in terms of, defaults to "v1".|
|value_from_field_ref_field_path|text|Path of the field to select in the specified API version.|
|value_from_resource_field_ref_container_name|text|Container name: required for volumes, optional for env vars|
|value_from_resource_field_ref_resource|text|Required: resource to select|
|value_from_resource_field_ref_divisor_format|text||
|value_from_config_map_key_ref_local_object_reference_name|text|Name of the referent.|
|value_from_config_map_key_ref_key|text|The key to select.|
|value_from_config_map_key_ref_optional|boolean|Specify whether the ConfigMap or its key must be defined|
|value_from_secret_key_ref_local_object_reference_name|text|Name of the referent.|
|value_from_secret_key_ref_key|text|The key of the secret to select from|
|value_from_secret_key_ref_optional|boolean|Specify whether the Secret or its key must be defined.|
