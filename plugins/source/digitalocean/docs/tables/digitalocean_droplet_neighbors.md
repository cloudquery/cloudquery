# Table: digitalocean_droplet_neighbors

The primary key for this table is **neighbor_id**.

## Relations

This table depends on [digitalocean_droplets](digitalocean_droplets.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|neighbor_id (PK)|Int|
|droplet_id|Int|