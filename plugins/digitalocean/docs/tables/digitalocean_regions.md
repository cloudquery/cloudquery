
# Table: digitalocean_regions
Region represents a DigitalOcean Region
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|slug|text|A human-readable string that is used as a unique identifier for each region.|
|name|text|The display name of the region.  This will be a full name that is used in the control panel and other interfaces.|
|sizes|text[]|This attribute is set to an array which contains the identifying slugs for the sizes available in this region.|
|available|boolean|This is a boolean value that represents whether new Droplets can be created in this region.|
|features|text[]|This attribute is set to an array which contains features available in this region|
