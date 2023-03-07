# Table: digitalocean_domain_records

The primary key for this table is **id**.

## Relations

This table depends on [digitalocean_domains](digitalocean_domains).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|type|String|
|name|String|
|data|String|
|priority|Int|
|port|Int|
|ttl|Int|
|weight|Int|
|flags|Int|
|tag|String|