# Table: digitalocean_firewalls



The primary key for this table is **id**.


## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_id|UUID|
|_cq_parent_id|UUID|
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|droplet_ids|IntArray|
|id (PK)|String|
|name|String|
|status|String|
|inbound_rules|JSON|
|outbound_rules|JSON|
|tags|StringArray|
|created_at|String|
|pending_changes|JSON|