
# Table: k8s_core_pod_ephemeral_container_statuses
ContainerStatus contains details for the current status of this container.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|pod_cq_id|uuid|Unique CloudQuery ID of k8s_core_pods table (FK)|
|name|text|This must be a DNS_LABEL|
|state|jsonb|Details about the container's current condition.|
|last_state|jsonb|Details about the container's last termination condition.|
|ready|boolean|Specifies whether the container has passed its readiness probe.|
|restart_count|integer|The number of times the container has been restarted.|
|image|text|The image the container is running.|
|image_id|text|ImageID of the container's image.|
|container_id|text|Container's ID in the format 'docker://<container_id>'.|
|started|boolean|Specifies whether the container has passed its startup probe.|
