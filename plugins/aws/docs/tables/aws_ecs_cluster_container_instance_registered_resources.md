
# Table: aws_ecs_cluster_container_instance_registered_resources
Describes the resources available for a container instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_container_instance_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_container_instances table (FK)|
|double_value|float|When the doubleValue type is set, the value of the resource must be a double precision floating-point type.|
|integer_value|integer|When the integerValue type is set, the value of the resource must be an integer.|
|long_value|bigint|When the longValue type is set, the value of the resource must be an extended precision floating-point type.|
|name|text|The name of the resource, such as CPU, MEMORY, PORTS, PORTS_UDP, or a user-defined resource.|
|string_set_value|text[]|When the stringSetValue type is set, the value of the resource must be a string type.|
|type|text|The type of the resource|
