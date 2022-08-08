
# Table: digitalocean_floating_ips
FloatingIP represents a Digital Ocean floating IP.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|droplet_id|bigint|Unique identifier of Droplet assigned the floating ip.|
|region_slug|text|A human-readable string that is used as a unique identifier for each region.|
|region_name|text|The display name of the region.  This will be a full name that is used in the control panel and other interfaces.|
|region_sizes|text[]|This attribute is set to an array which contains the identifying slugs for the sizes available in this region.|
|region_available|boolean|This is a boolean value that represents whether new Droplets can be created in this region.|
|region_features|text[]|This attribute is set to an array which contains features available in this region|
|ip|cidr|The public IP address of the floating IP. It also serves as its identifier.|
