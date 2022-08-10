
# Table: aws_ecs_cluster_container_instance_health_status_details
An object representing the result of a container instance health status check.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_container_instance_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_container_instances table (FK)|
|last_status_change|timestamp without time zone|The Unix timestamp for when the container instance health status last changed.|
|last_updated|timestamp without time zone|The Unix timestamp for when the container instance health status was last updated.|
|status|text|The container instance health status.|
|type|text|The type of container instance health status that was verified.|
