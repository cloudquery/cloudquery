
# Table: aws_ecs_cluster_service_load_balancers
The load balancer configuration to use with a service or task set
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_service_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_services table (FK)|
|container_name|text|The name of the container (as it appears in a container definition) to associate with the load balancer.|
|container_port|integer|The port on the container to associate with the load balancer|
|load_balancer_name|text|The name of the load balancer to associate with the Amazon ECS service or task set|
|target_group_arn|text|The full Amazon Resource Name (ARN) of the Elastic Load Balancing target group or groups associated with a service or task set|
