# Table: gcp_compute_interconnect_attachments

This table shows data for GCP Compute Interconnect Attachments.

https://cloud.google.com/compute/docs/reference/rest/v1/interconnectAttachments#InterconnectAttachment

The primary key for this table is **self_link**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|admin_enabled|`bool`|
|bandwidth|`utf8`|
|candidate_ipv6_subnets|`list<item: utf8, nullable>`|
|candidate_subnets|`list<item: utf8, nullable>`|
|cloud_router_ip_address|`utf8`|
|cloud_router_ipv6_address|`utf8`|
|cloud_router_ipv6_interface_id|`utf8`|
|configuration_constraints|`json`|
|creation_timestamp|`utf8`|
|customer_router_ip_address|`utf8`|
|customer_router_ipv6_address|`utf8`|
|customer_router_ipv6_interface_id|`utf8`|
|dataplane_version|`int64`|
|description|`utf8`|
|edge_availability_domain|`utf8`|
|encryption|`utf8`|
|google_reference_id|`utf8`|
|id|`int64`|
|interconnect|`utf8`|
|ipsec_internal_addresses|`list<item: utf8, nullable>`|
|kind|`utf8`|
|label_fingerprint|`utf8`|
|labels|`json`|
|mtu|`int64`|
|name|`utf8`|
|operational_status|`utf8`|
|pairing_key|`utf8`|
|partner_asn|`int64`|
|partner_metadata|`json`|
|private_interconnect_info|`json`|
|region|`utf8`|
|remote_service|`utf8`|
|router|`utf8`|
|satisfies_pzs|`bool`|
|self_link (PK)|`utf8`|
|stack_type|`utf8`|
|state|`utf8`|
|subnet_length|`int64`|
|type|`utf8`|
|vlan_tag8021q|`int64`|