# Table: gcp_kms_import_jobs

This table shows data for GCP Cloud Key Management Service (KMS) Import Jobs.

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.importJobs#ImportJob

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_keyrings](gcp_kms_keyrings).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|import_method|`utf8`|
|protection_level|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|generate_time|`timestamp[us, tz=UTC]`|
|expire_time|`timestamp[us, tz=UTC]`|
|expire_event_time|`timestamp[us, tz=UTC]`|
|state|`utf8`|
|public_key|`json`|
|attestation|`json`|