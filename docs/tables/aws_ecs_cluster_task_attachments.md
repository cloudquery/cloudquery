
# Table: aws_ecs_cluster_task_attachments
An object representing a container instance or task attachment.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_task_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_tasks table (FK)|
|details|jsonb|Details of the attachment|
|id|text|The unique identifier for the attachment.|
|status|text|The status of the attachment|
|type|text|The type of the attachment, such as ElasticNetworkInterface.|
