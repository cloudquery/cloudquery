# Table: gcp_aiplatform_dataset_locations

This table shows data for GCP AI Platform Dataset Locations.

https://cloud.google.com/api-gateway/docs/reference/rest/v1/projects.locations#Location

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_aiplatform_dataset_locations:
  - [gcp_aiplatform_datasets](gcp_aiplatform_datasets)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|location_id|`utf8`|
|display_name|`utf8`|
|labels|`json`|
|metadata|`json`|