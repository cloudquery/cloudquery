# Table: gcp_kms_locations

This table shows data for GCP Cloud Key Management Service (KMS) Locations.

https://cloud.google.com/kms/docs/reference/rest/v1/projects.locations.keyRings#KeyRing

The composite primary key for this table is (**project_id**, **name**).

## Relations

The following tables depend on gcp_kms_locations:
  - [gcp_kms_ekm_connections](gcp_kms_ekm_connections)
  - [gcp_kms_keyrings](gcp_kms_keyrings)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|location_id|`utf8`|
|display_name|`utf8`|
|labels|`json`|
|metadata|`json`|