# Table: gcp_compute_interconnects

This table shows data for GCP Compute Interconnects.

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|admin_enabled|`bool`|
|circuit_infos|`json`|
|creation_timestamp|`utf8`|
|customer_name|`utf8`|
|description|`utf8`|
|expected_outages|`json`|
|google_ip_address|`utf8`|
|google_reference_id|`utf8`|
|id|`int64`|
|interconnect_attachments|`list<item: utf8, nullable>`|
|interconnect_type|`utf8`|
|kind|`utf8`|
|link_type|`utf8`|
|location|`utf8`|
|name|`utf8`|
|noc_contact_email|`utf8`|
|operational_status|`utf8`|
|peer_ip_address|`utf8`|
|provisioned_link_count|`int64`|
|requested_link_count|`int64`|
|satisfies_pzs|`bool`|
|self_link (PK)|`utf8`|
|state|`utf8`|