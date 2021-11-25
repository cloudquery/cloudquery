
# Table: aws_ecs_cluster_container_instance_attributes
An attribute is a name-value pair associated with an Amazon ECS object. Attributes enable you to extend the Amazon ECS data model by adding custom metadata to your resources
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_container_instance_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_container_instances table (FK)|
|name|text|The name of the attribute|
|target_id|text|The ID of the target|
|target_type|text|The type of the target with which to attach the attribute|
|value|text|The value of the attribute|
