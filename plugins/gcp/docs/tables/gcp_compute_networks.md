
# Table: gcp_compute_networks
Represents a VPC Network resource  Networks connect resources to each other and to the internet
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|project_id|text|GCP Project Id of the resource|
|ip_v4_range|text|Deprecated in favor of subnet mode networks The range of internal addresses that are legal on this network This range is a CIDR specification, for example: 19216800/16 Provided by the client when the network is created|
|auto_create_subnetworks|boolean|Must be set to create a VPC network If not set, a legacy network is created  When set to true, the VPC network is created in auto mode When set to false, the VPC network is created in custom mode  An auto mode VPC network starts with one subnet per region Each subnet has a predetermined range as described in Auto mode VPC network IP ranges  For custom mode VPC networks, you can add subnets using the subnetworks insert method|
|creation_timestamp|text|Creation timestamp in RFC3339 text format|
|description|text|An optional description of this resource Provide this field when you create the resource|
|gateway_ip_v4|text|The gateway address for default routing out of the network, selected by GCP|
|resource_id|text|The unique identifier for the resource This identifier is defined by the server|
|kind|text|Type of the resource Always compute#network for networks|
|mtu|bigint|Maximum Transmission Unit in bytes The minimum value for this field is 1460 and the maximum value is 1500 bytes|
|name|text|Name of the resource Provided by the client when the resource is created|
|routing_config_routing_mode|text|The network-wide routing mode to use If set to REGIONAL, this network's Cloud Routers will only advertise routes with subnets of this network in the same region as the router If set to GLOBAL, this network's Cloud Routers will advertise routes with all subnets of this network, across regions|
|self_link|text|Server-defined URL for the resource|
|subnetworks|text[]|Server-defined fully-qualified URLs for all subnetworks in this VPC network|
