
# Table: aws_ecs_cluster_service_events
The details for an event that's associated with a service.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_service_cq_id|uuid|Unique CloudQuery ID of aws_ecs_cluster_services table (FK)|
|created_at|timestamp without time zone|The Unix timestamp for the time when the event was triggered.|
|id|text|The ID string for the event.|
|message|text|The event message.|
