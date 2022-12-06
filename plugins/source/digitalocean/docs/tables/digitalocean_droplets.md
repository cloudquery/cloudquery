# Table: digitalocean_droplets



The primary key for this table is **id**.

## Relations

The following tables depend on digitalocean_droplets:
  - [digitalocean_droplet_neighbors](digitalocean_droplet_neighbors.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|backup_ids|IntArray|
|snapshot_ids|IntArray|
|volume_ids|StringArray|
|id (PK)|Int|
|name|String|
|memory|Int|
|vcpus|Int|
|disk|Int|
|region|JSON|
|image|JSON|
|size|JSON|
|size_slug|String|
|next_backup_window|JSON|
|features|StringArray|
|locked|Bool|
|status|String|
|networks|JSON|
|created_at|String|
|kernel|JSON|
|tags|StringArray|
|vpc_uuid|String|