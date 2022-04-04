
# Table: azure_network_express_route_circuit_peerings
Peering in an ExpressRouteCircuit resource.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|express_route_circuit_cq_id|uuid|Unique CloudQuery ID of azure_network_express_route_circuits table (FK)|
|id|text|Resource ID.|
|azure_asn|integer|The Azure ASN.|
|etag|text|A unique read-only string that changes whenever the resource is updated.|
|express_route_connection_id|text|The ID of the ExpressRouteConnection.|
|gateway_manager_etag|text|The GatewayManager Etag.|
|ipv6_peering_config_microsoft_peering_config|jsonb|The Microsoft peering configuration.|
|ipv6_peering_config_primary_peer_address_prefix|text|The primary address prefix.|
|ipv6_peering_config_route_filter_id|text|The reference to the RouteFilter resource.|
|ipv6_peering_config_secondary_peer_address_prefix|text|The secondary address prefix.|
|ipv6_peering_config_state|text|The state of peering.|
|last_modified_by|text|Who was the last to modify the peering.|
|microsoft_peering_config_advertised_communities|text[]|The communities of bgp peering. Specified for microsoft peering.|
|microsoft_peering_config_advertised_public_prefixes|text[]|The reference to AdvertisedPublicPrefixes.|
|microsoft_peering_config_advertised_public_prefixes_state|text|The advertised public prefix state of the Peering resource.|
|microsoft_peering_config_customer_asn|integer|The CustomerASN of the peering.|
|microsoft_peering_config_legacy_mode|integer|The legacy mode of the peering.|
|microsoft_peering_config_routing_registry_name|text|The RoutingRegistryName of the configuration.|
|name|text|Resource name.|
|peer_asn|bigint|The peer ASN.|
|peering_type|text|The peering type.|
|primary_azure_port|text|The primary port.|
|primary_peer_address_prefix|text|The primary address prefix.|
|provisioning_state|text|The provisioning state of the express route circuit peering resource.|
|route_filter_id|text|The reference to the RouteFilter resource.|
|secondary_azure_port|text|The secondary port.|
|secondary_peer_address_prefix|text|The secondary address prefix.|
|shared_key|text|The shared key.|
|state|text|The peering state.|
|stats_primary_bytes_in|bigint|The Primary BytesIn of the peering.|
|stats_primary_bytes_out|bigint|The Primary BytesOut of the peering.|
|stats_secondary_bytes_in|bigint|The secondary BytesIn of the peering.|
|stats_secondary_bytes_out|bigint|The secondary BytesOut of the peering.|
|type|text|Resource type.|
|vlan_id|integer|The VLAN ID.|
