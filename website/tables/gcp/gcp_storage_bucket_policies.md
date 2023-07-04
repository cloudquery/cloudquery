# Table: gcp_storage_bucket_policies

This table shows data for GCP Storage Bucket Policies.

https://cloud.google.com/iam/docs/reference/rest/v1/Policy

The primary key for this table is **_cq_id**.

## Relations

This table depends on [gcp_storage_buckets](gcp_storage_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|bucket_name|`utf8`|
|bindings|`json`|