# Table: oracle_virtual_network_fast_connect_provider_services

This table shows data for Oracle Virtual Network Fast Connect Provider Services.

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
|id (PK)|`utf8`|
|private_peering_bgp_management|`utf8`|
|provider_name|`utf8`|
|provider_service_name|`utf8`|
|public_peering_bgp_management|`utf8`|
|customer_asn_management|`utf8`|
|provider_service_key_management|`utf8`|
|bandwith_shape_management|`utf8`|
|required_total_cross_connects|`int64`|
|type|`utf8`|
|description|`utf8`|
|supported_virtual_circuit_types|`list<item: utf8, nullable>`|