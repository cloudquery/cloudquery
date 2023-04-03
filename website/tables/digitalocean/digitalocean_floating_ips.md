# Table: digitalocean_floating_ips

This table shows data for DigitalOcean Floating IPs.

The primary key for this table is **ip**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|ip (PK)|String|
|region|JSON|
|droplet|JSON|
|project_id|String|
|locked|Bool|