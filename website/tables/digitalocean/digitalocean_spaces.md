# Table: digitalocean_spaces

This table shows data for DigitalOcean Spaces.

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on digitalocean_spaces:
  - [digitalocean_space_cors](digitalocean_space_cors)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|acls|`json`|
|bucket|`json`|
|location|`utf8`|
|public|`bool`|