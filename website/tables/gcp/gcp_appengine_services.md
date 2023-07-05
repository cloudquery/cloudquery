# Table: gcp_appengine_services

This table shows data for GCP App Engine Services.

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services#Service

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_appengine_services:
  - [gcp_appengine_versions](gcp_appengine_versions)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|split|`json`|
|labels|`json`|
|network_settings|`json`|