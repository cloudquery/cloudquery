# Table: gcp_kms_keyrings

This table shows data for GCP Cloud Key Management Service (KMS) Keyrings.

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_locations](gcp_kms_locations).

The following tables depend on gcp_kms_keyrings:
  - [gcp_kms_crypto_keys](gcp_kms_crypto_keys)
  - [gcp_kms_import_jobs](gcp_kms_import_jobs)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|