# Table: gcp_certificatemanager_certificate_issuance_configs

This table shows data for GCP Certificatemanager Certificate Issuance Configs.

https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificateIssuanceConfigs#CertificateIssuanceConfig

The composite primary key for this table is (**project_id**, **name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name (PK)|`utf8`|
|create_time|`timestamp[us, tz=UTC]`|
|update_time|`timestamp[us, tz=UTC]`|
|labels|`json`|
|description|`utf8`|
|certificate_authority_config|`json`|
|lifetime|`int64`|
|rotation_window_percentage|`int64`|
|key_algorithm|`utf8`|