
# Table: digitalocean_vpcs
VPC represents a DigitalOcean Virtual Private Cloud configuration.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|A unique ID that can be used to identify and reference the VPC.|
|urn|text|The uniform resource name (URN) for the resource in the format do:resource_type:resource_id.|
|name|text|The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods.|
|description|text|A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters.|
|ip_range|cidr|The range of IP addresses in the VPC in CIDR notation. Network ranges cannot overlap with other networks in the same account and must be in range of private addresses as defined in RFC1918. It may not be smaller than `/24` nor larger than `/16`. If no IP range is specified, a `/20` network range is generated that won't conflict with other VPC networks in your account.|
|region_slug|text|The slug identifier for the region where the VPC will be created.|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format.|
|default|boolean|A boolean value indicating whether or not the VPC is the default network for the region. All applicable resources are placed into the default VPC network unless otherwise specified during their creation. The `default` field cannot be unset from `true`. If you want to set a new default VPC network, update the `default` field of another VPC network in the same region. The previous network's `default` field will be set to `false` when a new default VPC has been defined.|
