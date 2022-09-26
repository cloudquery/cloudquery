# Table: gcp_kms_keyrings


The primary key for this table is **_cq_id**.

## Relations
The following tables depend on `gcp_kms_keyrings`:
  - [`gcp_kms_crypto_keys`](gcp_kms_crypto_keys.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|name|String|
|create_time|JSON|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|