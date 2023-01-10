# Table: gcp_kms_keyrings

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_kms_keyrings:
  - [gcp_kms_crypto_keys](gcp_kms_crypto_keys.md)

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