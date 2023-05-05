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
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|Int|
|name|String|
|memory|Int|
|vcpus|Int|
|disk|Int|
|region|JSON|
|image|JSON|
|size|JSON|
|size_slug|String|
|backup_ids|IntArray|
|next_backup_window|JSON|
|snapshot_ids|IntArray|
|features|StringArray|
|locked|Bool|
|status|String|
|networks|JSON|
|created_at|String|
|kernel|JSON|
|tags|StringArray|
|volume_ids|StringArray|
|vpc_uuid|String|