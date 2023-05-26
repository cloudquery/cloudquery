# Table: digitalocean_floating_ips

This table shows data for DigitalOcean Floating IPs.

Deprecated. https://docs.digitalocean.com/reference/api/api-reference/#tag/Floating-IPs

The primary key for this table is **ip**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|ip (PK)|utf8|
|region|json|
|droplet|json|
|project_id|utf8|
|locked|bool|