# Table: gcp_kms_locations

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_kms_locations:
  - [gcp_kms_ekm_connections](gcp_kms_ekm_connections.md)
  - [gcp_kms_keyrings](gcp_kms_keyrings.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|location_id|String|
|display_name|String|
|labels|JSON|
|metadata|JSON|