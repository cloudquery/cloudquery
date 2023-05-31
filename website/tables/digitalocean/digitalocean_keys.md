# Table: digitalocean_keys

This table shows data for DigitalOcean Keys.

https://docs.digitalocean.com/reference/api/api-reference/#tag/SSH-Keys

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|name|`utf8`|
|fingerprint|`utf8`|
|public_key|`utf8`|