
# Table: gcp_compute_addresses
Addresses for GFE-based external HTTP(S) load balancers.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|address|text|The static IP address represented by this resource|
|address_type|text|The type of address to reserve, either INTERNAL or EXTERNAL If unspecified, defaults to EXTERNAL|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource Provide this field when you create the resource|
|address_id|text|The unique identifier for the resource This identifier is defined by the server|
|ip_version|text|The IP version that will be used by this address Valid options are IPV4 or IPV6 This can only be specified for a global address|
|kind|text|Type of the resource Always compute#address for addresses|
|name|text|Name of the resource Provided by the client when the resource is created|
|network|text|The URL of the network in which to reserve the address This field can only be used with INTERNAL type with the VPC_PEERING purpose|
|network_tier|text|This signifies the networking tier used for configuring this address|
|prefix_length|bigint|The prefix length if the resource represents an IP range|
|purpose|text|The purpose of this resource, which can be one of the following values: - `GCE_ENDPOINT` for addresses that are used by VM instances, alias IP ranges, internal load balancers, and similar resources - `DNS_RESOLVER` for a DNS resolver address in a subnetwork - `VPC_PEERING` for addresses that are reserved for VPC peer networks - `NAT_AUTO` for addresses that are external IP addresses automatically reserved for Cloud NAT - `IPSEC_INTERCONNECT` for addresses created from a private IP range that are reserved for a VLAN attachment in an IPsec-encrypted Cloud Interconnect configuration These addresses are regional resources|
|region|text|The URL of the region where a regional address resides For regional addresses, you must specify the region as a path parameter in the HTTP request URL This field is not applicable to global addresses|
|self_link|text|Server-defined URL for the resource|
|status|text|The status of the address, which can be one of RESERVING, RESERVED, or IN_USE An address that is RESERVING is currently in the process of being reserved A RESERVED address is currently reserved and available to use An IN_USE address is currently being used by another resource and is not available|
|subnetwork|text|The URL of the subnetwork in which to reserve the address If an IP address is specified, it must be within the subnetwork's IP range This field can only be used with INTERNAL type with a GCE_ENDPOINT or DNS_RESOLVER purpose|
|users|text[]|The URLs of the resources that are using this address|
