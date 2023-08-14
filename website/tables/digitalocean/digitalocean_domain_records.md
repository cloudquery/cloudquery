# Table: digitalocean_domain_records

This table shows data for DigitalOcean Domain Records.

https://docs.digitalocean.com/reference/api/api-reference/#operation/domains_list_records

The primary key for this table is **id**.

## Relations

This table depends on [digitalocean_domains](digitalocean_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|type|`utf8`|
|name|`utf8`|
|data|`utf8`|
|priority|`int64`|
|port|`int64`|
|ttl|`int64`|
|weight|`int64`|
|flags|`int64`|
|tag|`utf8`|