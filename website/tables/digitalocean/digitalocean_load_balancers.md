# Table: digitalocean_load_balancers

This table shows data for DigitalOcean Load Balancers.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Load-Balancers

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|name|String|
|ip|String|
|size|String|
|size_unit|Int|
|algorithm|String|
|status|String|
|created_at|String|
|forwarding_rules|JSON|
|health_check|JSON|
|sticky_sessions|JSON|
|region|JSON|
|droplet_ids|IntArray|
|tag|String|
|tags|StringArray|
|redirect_http_to_https|Bool|
|enable_proxy_protocol|Bool|
|enable_backend_keepalive|Bool|
|vpc_uuid|String|
|disable_lets_encrypt_dns_records|Bool|
|validate_only|Bool|
|project_id|String|
|http_idle_timeout_seconds|Int|
|firewall|JSON|