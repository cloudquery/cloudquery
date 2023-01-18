# Table: gcp_appengine_services

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services#Service

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_appengine_services:
  - [gcp_appengine_versions](gcp_appengine_versions.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|id|String|
|split|JSON|
|labels|JSON|
|network_settings|JSON|