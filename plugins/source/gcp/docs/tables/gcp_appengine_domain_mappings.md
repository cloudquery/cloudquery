# Table: gcp_appengine_domain_mappings

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.domainMappings#DomainMapping

The composite primary key for this table is (**project_id**, **name**).

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
|ssl_settings|JSON|
|resource_records|JSON|