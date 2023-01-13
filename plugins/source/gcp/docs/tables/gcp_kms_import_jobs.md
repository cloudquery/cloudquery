# Table: gcp_kms_import_jobs

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings.importJobs#ImportJob

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_keyrings](gcp_kms_keyrings.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name (PK)|String|
|import_method|String|
|protection_level|String|
|create_time|Timestamp|
|generate_time|Timestamp|
|expire_time|Timestamp|
|expire_event_time|Timestamp|
|state|String|
|public_key|JSON|
|attestation|JSON|