# Table: gcp_appengine_domain_mappings

This table shows data for GCP App Engine Domain Mappings.

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.domainMappings#DomainMapping

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|ssl_settings|`json`|
|resource_records|`json`|