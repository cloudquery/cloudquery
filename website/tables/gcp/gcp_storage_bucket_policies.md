# Table: gcp_storage_bucket_policies

This table shows data for GCP Storage Bucket Policies.

https://cloud.google.com/iam/docs/reference/rest/v1/Policy

The composite primary key for this table is ().

## Relations

This table depends on [gcp_storage_buckets](gcp_storage_buckets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|project_id|`utf8`|
|bucket_name|`utf8`|
|bindings|`json`|