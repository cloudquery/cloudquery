# Table: gcp_domains_registrations

This table shows data for GCP Domains Registrations.

https://cloud.google.com/domains/docs/reference/rest/v1beta1/projects.locations.registrations#Registration

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|domain_name|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|expire_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|issues|`list<item: int64, nullable>`|
|labels|`json`|
|management_settings|`json`|
|dns_settings|`json`|
|contact_settings|`json`|
|pending_contact_settings|`json`|
|supported_privacy|`list<item: int64, nullable>`|