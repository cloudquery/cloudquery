# Table: digitalocean_load_balancers

This table shows data for DigitalOcean Load Balancers.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Load-Balancers

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|ip|`utf8`|
|size|`utf8`|
|size_unit|`int64`|
|algorithm|`utf8`|
|status|`utf8`|
|created_at|`utf8`|
|forwarding_rules|`json`|
|health_check|`json`|
|sticky_sessions|`json`|
|region|`json`|
|droplet_ids|`list<item: int64, nullable>`|
|tag|`utf8`|
|tags|`list<item: utf8, nullable>`|
|redirect_http_to_https|`bool`|
|enable_proxy_protocol|`bool`|
|enable_backend_keepalive|`bool`|
|vpc_uuid|`utf8`|
|disable_lets_encrypt_dns_records|`bool`|
|validate_only|`bool`|
|project_id|`utf8`|
|http_idle_timeout_seconds|`int64`|
|firewall|`json`|