# Table: gcp_certificatemanager_certificate_map_entries

https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificateMaps.certificateMapEntries#CertificateMapEntry

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_certificatemanager_certificate_maps](gcp_certificatemanager_certificate_maps.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|description|String|
|create_time|Timestamp|
|update_time|Timestamp|
|labels|JSON|
|certificates|StringArray|
|state|String|