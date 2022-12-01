# Table: gcp_compute_interconnects



The primary key for this table is **self_link**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|String|
|self_link (PK)|String|
|admin_enabled|Bool|
|circuit_infos|JSON|
|creation_timestamp|String|
|customer_name|String|
|description|String|
|expected_outages|JSON|
|google_ip_address|String|
|google_reference_id|String|
|id|Int|
|interconnect_attachments|StringArray|
|interconnect_type|String|
|kind|String|
|link_type|String|
|location|String|
|name|String|
|noc_contact_email|String|
|operational_status|String|
|peer_ip_address|String|
|provisioned_link_count|Int|
|requested_link_count|Int|
|satisfies_pzs|Bool|
|state|String|