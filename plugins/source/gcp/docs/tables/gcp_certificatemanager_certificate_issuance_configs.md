# Table: gcp_certificatemanager_certificate_issuance_configs

https://cloud.google.com/certificate-manager/docs/reference/rest/v1/projects.locations.certificateIssuanceConfigs#CertificateIssuanceConfig

The composite primary key for this table is (**project_id**, **name**).

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
|update_time|Timestamp|
|labels|JSON|
|description|String|
|certificate_authority_config|JSON|
|lifetime|Int|
|rotation_window_percentage|Int|
|key_algorithm|String|