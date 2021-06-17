
# Table: gcp_compute_instance_network_interface_access_configs
An access configuration attached to an instance's network interface Only one access config per instance is supported
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_network_interface_id|uuid|Unique ID of gcp_compute_instance_network_interfaces table (FK)|
|kind|text|Type of the resource Always compute#accessConfig for access configs|
|name|text|The name of this access configuration The default and recommended name is External NAT, but you can use any arbitrary string, such as My external IP or Network Access|
|nat_ip|text|An external IP address associated with this instance Specify an unused static external IP address available to the project or leave this field undefined to use an IP from a shared ephemeral IP address pool If you specify a static external IP address, it must live in the same region as the zone of the instance|
|network_tier|text|This signifies the networking tier used for configuring this access configuration and can only take the following values: PREMIUM, STANDARD  If an AccessConfig is specified without a valid external IP address, an ephemeral IP will be created with this networkTier  If an AccessConfig with a valid external IP address is specified, it must match that of the networkTier associated with the Address resource owning that IP|
|public_ptr_domain_name|text|The DNS domain name for the public PTR record You can set this field only if the `setPublicPtr` field is enabled|
|set_public_ptr|boolean|Specifies whether a public DNS 'PTR' record should be created to map the external IP address of the instance to a DNS domain name|
|type|text|The type of configuration The default and only option is ONE_TO_ONE_NAT|
