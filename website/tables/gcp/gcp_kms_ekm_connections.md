# Table: gcp_kms_ekm_connections

This table shows data for GCP Cloud Key Management Service (KMS) Cloud External Key Manager (EKM) Connections.

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.ekmConnections#EkmConnection

The composite primary key for this table is (**project_id**, **name**).

## Relations

This table depends on [gcp_kms_locations](gcp_kms_locations).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|service_resolvers|`json`|
|etag|`utf8`|
|key_management_mode|`utf8`|
|crypto_space_path|`utf8`|