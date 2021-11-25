
# Table: aws_ecs_cluster_service_events
Details on an event associated with a service.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_service_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_services table (FK)|
|created_at|timestamp without time zone|The Unix timestamp for when the event was triggered.|
|id|text|The ID string of the event.|
|message|text|The event message.|
