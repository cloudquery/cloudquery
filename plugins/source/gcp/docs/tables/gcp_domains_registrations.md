# Table: gcp_domains_registrations



The primary key for this table is **_cq_id**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|name|String|
|domain_name|String|
|create_time|Timestamp|
|expire_time|Timestamp|
|state|String|
|issues|IntArray|
|labels|JSON|
|management_settings|JSON|
|dns_settings|JSON|
|contact_settings|JSON|
|pending_contact_settings|JSON|
|supported_privacy|IntArray|