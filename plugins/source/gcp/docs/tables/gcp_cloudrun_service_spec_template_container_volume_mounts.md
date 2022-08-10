
# Table: gcp_cloudrun_service_spec_template_container_volume_mounts
VolumeMount describes a mounting of a Volume within a container
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_spec_template_container_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_service_spec_template_containers table (FK)|
|mount_path|text|Path within the container at which the volume should be mounted|
|name|text|The name of the volume|
|read_only|boolean|Only true is accepted|
|sub_path|text|Path within the volume from which the container's volume should be mounted|
