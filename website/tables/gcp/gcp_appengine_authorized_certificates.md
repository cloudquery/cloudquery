# Table: gcp_appengine_authorized_certificates

This table shows data for GCP App Engine Authorized Certificates.

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.authorizedCertificates#AuthorizedCertificate

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|id|`utf8`|
|display_name|`utf8`|
|domain_names|`list<item: utf8, nullable>`|
|expire_time|`timestamp[us, tz=UTC]`|
|certificate_raw_data|`json`|
|managed_certificate|`json`|
|visible_domain_mappings|`list<item: utf8, nullable>`|
|domain_mappings_count|`int64`|