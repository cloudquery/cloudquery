# Table: digitalocean_droplets

This table shows data for DigitalOcean Droplets.

https://docs.digitalocean.com/reference/api/api-reference/#operation/droplets_list

The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_droplets:
  - [digitalocean_droplet_neighbors](digitalocean_droplet_neighbors)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|id (PK)|`int64`|
|name|`utf8`|
|memory|`int64`|
|vcpus|`int64`|
|disk|`int64`|
|region|`json`|
|image|`json`|
|size|`json`|
|size_slug|`utf8`|
|backup_ids|`list<item: int64, nullable>`|
|next_backup_window|`json`|
|snapshot_ids|`list<item: int64, nullable>`|
|features|`list<item: utf8, nullable>`|
|locked|`bool`|
|status|`utf8`|
|networks|`json`|
|created_at|`utf8`|
|kernel|`json`|
|tags|`list<item: utf8, nullable>`|
|volume_ids|`list<item: utf8, nullable>`|
|vpc_uuid|`utf8`|