# Table: gcp_domains_registrations


The primary key for this table is **_cq_id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|project_id|String|
|name|String|
|domain_name|String|
|create_time|JSON|
|expire_time|JSON|
|state|Int|
|issues|IntArray|
|labels|JSON|
|management_settings|JSON|
|dns_settings|JSON|
|contact_settings|JSON|
|pending_contact_settings|JSON|
|supported_privacy|IntArray|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|