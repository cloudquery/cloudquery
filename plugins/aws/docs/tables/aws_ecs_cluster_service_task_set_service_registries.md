
# Table: aws_ecs_cluster_service_task_set_service_registries
The details for the service registry
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_service_task_set_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_service_task_sets table (FK)|
|container_name|text|The container name value to be used for your service discovery service|
|container_port|integer|The port value to be used for your service discovery service|
|port|integer|The port value used if your service discovery service specified an SRV record. This field might be used if both the awsvpc network mode and SRV records are used.|
|arn|text|The Amazon Resource Name (ARN) of the service registry|
