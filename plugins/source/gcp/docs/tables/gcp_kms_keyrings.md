# Table: gcp_kms_keyrings

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_locations](gcp_kms_locations.md).

The following tables depend on gcp_kms_keyrings:
  - [gcp_kms_crypto_keys](gcp_kms_crypto_keys.md)
  - [gcp_kms_import_jobs](gcp_kms_import_jobs.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|create_time|Timestamp|