
# Table: gcp_compute_instance_network_interface_alias_ip_ranges
An alias IP range attached to an instance's network interface
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_network_interface_cq_id|uuid|Unique ID of gcp_compute_instance_network_interfaces table (FK)|
|instance_network_interface_name|text||
|ip_cidr_range|text|The IP alias ranges to allocate for this interface This IP CIDR range must belong to the specified subnetwork and cannot contain IP addresses reserved by system or used by other network interfaces This range may be a single IP address (such as 10234), a netmask (such as /24) or a CIDR-formatted string (such as 10120/24)|
|subnetwork_range_name|text|The name of a subnetwork secondary IP range from which to allocate an IP alias range If not specified, the primary range of the subnetwork is used|
