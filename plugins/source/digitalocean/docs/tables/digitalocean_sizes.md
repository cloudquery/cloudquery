
# Table: digitalocean_sizes
Size represents a DigitalOcean Size
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|slug|text|A human-readable string that is used to uniquely identify each size.|
|memory|bigint|The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes.|
|vcpus|bigint|The integer of number CPUs allocated to Droplets of this size.|
|disk|bigint|The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes.|
|price_monthly|float|This attribute describes the monthly cost of this Droplet size if the Droplet is kept for an entire month. The value is measured in US dollars.|
|price_hourly|float|This describes the price of the Droplet size as measured hourly. The value is measured in US dollars.|
|regions|text[]|An array containing the region slugs where this size is available for Droplet creates.|
|available|boolean|This is a boolean value that represents whether new Droplets can be created with this size.|
|transfer|float|The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes.|
|description|text|A string describing the class of Droplets created from this size. For example: Basic, General Purpose, CPU-Optimized, Memory-Optimized, or Storage-Optimized.|
