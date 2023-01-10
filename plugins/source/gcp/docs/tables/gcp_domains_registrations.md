# Table: gcp_domains_registrations

https://cloud.google.com/domains/docs/reference/rest/v1beta1/projects.locations.registrations#Registration

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
|domain_name|String|
|create_time|Timestamp|
|expire_time|Timestamp|
|state|String|
|issues|IntArray|
|labels|JSON|
|management_settings|JSON|
|dns_settings|JSON|
|contact_settings|JSON|
|pending_contact_settings|JSON|
|supported_privacy|IntArray|