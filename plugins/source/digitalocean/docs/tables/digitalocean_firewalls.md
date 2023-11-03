# Table: digitalocean_firewalls

This table shows data for DigitalOcean Firewalls.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Firewalls

The primary key for this table is **id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`utf8`|
|name|`utf8`|
|status|`utf8`|
|inbound_rules|`json`|
|outbound_rules|`json`|
|droplet_ids|`list<item: int64, nullable>`|
|tags|`list<item: utf8, nullable>`|
|created_at|`utf8`|
|pending_changes|`json`|