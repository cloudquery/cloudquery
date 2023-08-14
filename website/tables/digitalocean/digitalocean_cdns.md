# Table: digitalocean_cdns

This table shows data for DigitalOcean CDNs.

https://docs.digitalocean.com/reference/api/api-reference/#tag/CDN-Endpoints

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|origin|`utf8`|
|endpoint|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|ttl|`int64`|
|certificate_id|`utf8`|
|custom_domain|`utf8`|