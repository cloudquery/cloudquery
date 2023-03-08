# Table: oracle_virtualnetwork_virtual_circuits

The composite primary key for this table is (**region**, **compartment_id**, **id**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|region (PK)|String|
|compartment_id (PK)|String|
|id (PK)|String|
|bandwidth_shape_name|String|
|bgp_management|String|
|bgp_session_state|String|
|bgp_ipv6_session_state|String|
|cross_connect_mappings|JSON|
|routing_policy|StringArray|
|bgp_admin_state|String|
|is_bfd_enabled|Bool|
|customer_bgp_asn|Int|
|customer_asn|Int|
|defined_tags|JSON|
|display_name|String|
|freeform_tags|JSON|
|gateway_id|String|
|lifecycle_state|String|
|oracle_bgp_asn|Int|
|provider_name|String|
|provider_service_id|String|
|provider_service_key_name|String|
|provider_service_name|String|
|provider_state|String|
|public_prefixes|StringArray|
|reference_comment|String|
|service_type|String|
|time_created|Timestamp|
|type|String|
|ip_mtu|String|