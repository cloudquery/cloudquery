# Table: gcp_certificatemanager_certificates

This table shows data for GCP Certificatemanager Certificates.

https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificates#Certificate

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|description|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|san_dnsnames|`list<item: utf8, nullable>`|
|pem_certificate|`utf8`|
|expire_time|`timestamp[us, tz=UTC]`|
|scope|`utf8`|