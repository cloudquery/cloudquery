
# Table: k8s_core_pod_ephemeral_container_volume_mounts
VolumeMount describes a mounting of a Volume within a container.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|pod_ephemeral_container_cq_id|uuid|Unique CloudQuery ID of k8s_core_pod_init_containers table (FK)|
|name|text|This must match the Name of a Volume.|
|read_only|boolean|Mounted read-only if true, read-write otherwise (false or unspecified). Defaults to false. +optional|
|mount_path|text|Path within the container at which the volume should be mounted|
|sub_path|text|Path within the volume from which the container's volume should be mounted. Defaults to "" (volume's root). +optional|
|mount_propagation|text|mountPropagation determines how mounts are propagated from the host to container and the other way around. When not set, MountPropagationNone is used. This field is beta in 1.10. +optional|
|sub_path_expr|text|Expanded path within the volume from which the container's volume should be mounted. Behaves similarly to SubPath but environment variable references $(VAR_NAME) are expanded using the container's environment. Defaults to "" (volume's root). SubPathExpr and SubPath are mutually exclusive. +optional|
