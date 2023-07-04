# Table: oracle_virtualnetwork_virtual_circuits

This table shows data for Oracle Virtual Network Virtual Circuits.

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|region (PK)|`utf8`|
|compartment_id (PK)|`utf8`|
|bandwidth_shape_name|`utf8`|
|bgp_management|`utf8`|
|bgp_session_state|`utf8`|
|bgp_ipv6_session_state|`utf8`|
|cross_connect_mappings|`json`|
|routing_policy|`list<item: utf8, nullable>`|
|bgp_admin_state|`utf8`|
|is_bfd_enabled|`bool`|
|customer_bgp_asn|`int64`|
|customer_asn|`int64`|
|defined_tags|`json`|
|display_name|`utf8`|
|freeform_tags|`json`|
|gateway_id|`utf8`|
|id (PK)|`utf8`|
|lifecycle_state|`utf8`|
|oracle_bgp_asn|`int64`|
|provider_name|`utf8`|
|provider_service_id|`utf8`|
|provider_service_key_name|`utf8`|
|provider_service_name|`utf8`|
|provider_state|`utf8`|
|public_prefixes|`list<item: utf8, nullable>`|
|reference_comment|`utf8`|
|service_type|`utf8`|
|time_created|`timestamp[us, tz=UTC]`|
|type|`utf8`|
|ip_mtu|`utf8`|