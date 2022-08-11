
# Table: gcp_compute_network_peerings
A network peering attached to a network resource The message includes the peering name, peer network, peering state, and a flag indicating whether Google Compute Engine should automatically create routes for the peering
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|network_cq_id|uuid|Unique ID of gcp_compute_networks table (FK)|
|network_name|text||
|auto_create_routes|boolean|This field will be deprecated soon Use the exchange_subnet_routes field instead Indicates whether full mesh connectivity is created and managed automatically between peered networks Currently this field should always be true since Google Compute Engine will automatically create and manage subnetwork routes between two networks when peering state is ACTIVE|
|exchange_subnet_routes|boolean|Indicates whether full mesh connectivity is created and managed automatically between peered networks Currently this field should always be true since Google Compute Engine will automatically create and manage subnetwork routes between two networks when peering state is ACTIVE|
|export_custom_routes|boolean|Whether to export the custom routes to peer network|
|export_subnet_routes_with_public_ip|boolean|Whether subnet routes with public IP range are exported The default value is true, all subnet routes are exported The IPv4 special-use ranges (https://enwikipediaorg/wiki/IPv4#Special_addresses) are always exported to peers and are not controlled by this field|
|import_custom_routes|boolean|Whether to import the custom routes from peer network|
|import_subnet_routes_with_public_ip|boolean|Whether subnet routes with public IP range are imported The default value is false The IPv4 special-use ranges (https://enwikipediaorg/wiki/IPv4#Special_addresses) are always imported from peers and are not controlled by this field|
|name|text|Name of this peering Provided by the client when the peering is created The name must comply with RFC1035 Specifically, the name must be 1-63 characters long and match regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` The first character must be a lowercase letter, and all the following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash|
|network|text|The URL of the peer network It can be either full URL or partial URL The peer network may belong to a different project If the partial URL does not contain project, it is assumed that the peer network is in the same project as the current network|
|peer_mtu|bigint|Maximum Transmission Unit in bytes|
|state|text|State for the peering, either `ACTIVE` or `INACTIVE` The peering is `ACTIVE` when there's a matching configuration in the peer network|
|state_details|text|Details about the current state of the peering|
