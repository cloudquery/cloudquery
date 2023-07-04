# Table: gcp_clouddeploy_delivery_pipelines

This table shows data for GCP Clouddeploy Delivery Pipelines.

https://cloud.google.com/deploy/docs/api/reference/rest/v1/projects.locations.deliveryPipelines#DeliveryPipeline

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_clouddeploy_delivery_pipelines:
  - [gcp_clouddeploy_releases](gcp_clouddeploy_releases)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|uid|`utf8`|
|description|`utf8`|
|annotations|`json`|
|labels|`json`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|condition|`json`|
|etag|`utf8`|
|suspended|`bool`|