# Table: digitalocean_regions

This table shows data for DigitalOcean Regions.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Regions

The primary key for this table is **slug**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|slug (PK)|`utf8`|
|name|`utf8`|
|sizes|`list<item: utf8, nullable>`|
|available|`bool`|
|features|`list<item: utf8, nullable>`|