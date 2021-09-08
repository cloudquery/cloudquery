
# Table: digitalocean_volumes

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|id|text|The unique identifier for the block storage volume.|
|region_slug|text|A human-readable string that is used as a unique identifier for each region.|
|region_name|text|The display name of the region.  This will be a full name that is used in the control panel and other interfaces.|
|region_sizes|text[]|This attribute is set to an array which contains the identifying slugs for the sizes available in this region.|
|region_available|boolean|This is a boolean value that represents whether new Droplets can be created in this region.|
|region_features|text[]|This attribute is set to an array which contains features available in this region|
|name|text|A human-readable name for the block storage volume. Must be lowercase and be composed only of numbers, letters and "-", up to a limit of 64 characters. The name must begin with a letter.|
|size_giga_bytes|bigint|The size of the block storage volume in GiB (1024^3).|
|description|text|An optional free-form text field to describe a block storage volume.|
|droplet_ids|integer[]|An array containing the IDs of the Droplets the volume is attached to. Note that at this time, a volume can only be attached to a single Droplet.|
|created_at|timestamp without time zone|A time value given in ISO8601 combined date and time format that represents when the block storage volume was created.|
|filesystem_type|text|The type of filesystem currently in-use on the volume.|
|filesystem_label|text|The label currently applied to the filesystem.|
|tags|text[]|A flat array of tag names as strings to be applied to the resource. Tag names may be for either existing or new tags.|
