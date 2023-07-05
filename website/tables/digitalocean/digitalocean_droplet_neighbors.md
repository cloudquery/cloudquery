# Table: digitalocean_droplet_neighbors

This table shows data for DigitalOcean Droplet Neighbors.

The primary key for this table is **neighbor_id**.

## Relations

This table depends on [digitalocean_droplets](digitalocean_droplets).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|neighbor_id (PK)|`int64`|
|droplet_id|`int64`|