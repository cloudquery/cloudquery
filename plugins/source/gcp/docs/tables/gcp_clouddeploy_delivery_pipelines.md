# Table: gcp_clouddeploy_delivery_pipelines

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines#DeliveryPipeline

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_clouddeploy_delivery_pipelines:
  - [gcp_clouddeploy_releases](gcp_clouddeploy_releases.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|uid|String|
|description|String|
|annotations|JSON|
|labels|JSON|
|create_time|Timestamp|
|update_time|Timestamp|
|condition|JSON|
|etag|String|
|suspended|Bool|