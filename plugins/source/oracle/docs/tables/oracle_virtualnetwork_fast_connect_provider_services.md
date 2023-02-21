# Table: oracle_virtualnetwork_fast_connect_provider_services

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
|private_peering_bgp_management|String|
|provider_name|String|
|provider_service_name|String|
|public_peering_bgp_management|String|
|customer_asn_management|String|
|provider_service_key_management|String|
|bandwith_shape_management|String|
|required_total_cross_connects|Int|
|type|String|
|description|String|
|supported_virtual_circuit_types|StringArray|