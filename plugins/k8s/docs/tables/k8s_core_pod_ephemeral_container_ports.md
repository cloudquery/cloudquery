
# Table: k8s_core_pod_ephemeral_container_ports
ContainerPort represents a network port in a single container.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|pod_ephemeral_container_cq_id|uuid|Unique CloudQuery ID of k8s_core_pod_init_containers table (FK)|
|name|text|If specified, this must be an IANA_SVC_NAME and unique within the pod|
|host_port|integer|Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this. +optional|
|container_port|integer|Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.|
|protocol|text|Protocol for port|
|host_ip|text|What host IP to bind the external port to. +optional|
