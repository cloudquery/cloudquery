# Table: digitalocean_sizes

This table shows data for DigitalOcean Sizes.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Sizes

The primary key for this table is **slug**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|slug (PK)|`utf8`|
|memory|`int64`|
|vcpus|`int64`|
|disk|`int64`|
|price_monthly|`float64`|
|price_hourly|`float64`|
|regions|`list<item: utf8, nullable>`|
|available|`bool`|
|transfer|`float64`|
|description|`utf8`|