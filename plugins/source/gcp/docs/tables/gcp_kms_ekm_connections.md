# Table: gcp_kms_ekm_connections

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.ekmConnections#EkmConnection

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_locations](gcp_kms_locations.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|create_time|Timestamp|
|service_resolvers|JSON|
|etag|String|