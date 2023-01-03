# Table: gcp_appengine_authorized_certificates

https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.authorizedCertificates#AuthorizedCertificate

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
|id|String|
|display_name|String|
|domain_names|StringArray|
|expire_time|Timestamp|
|certificate_raw_data|JSON|
|managed_certificate|JSON|
|visible_domain_mappings|StringArray|
|domain_mappings_count|Int|