# Table: gcp_certificatemanager_certificate_map_entries

This table shows data for GCP Certificatemanager Certificate Map Entries.

https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificateMaps.certificateMapEntries#CertificateMapEntry

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_certificatemanager_certificate_maps](gcp_certificatemanager_certificate_maps).

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
|certificates|`list<item: utf8, nullable>`|
|state|`utf8`|