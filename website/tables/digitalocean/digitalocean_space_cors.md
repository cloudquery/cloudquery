# Table: digitalocean_space_cors

This table shows data for DigitalOcean Space CORS.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [digitalocean_spaces](digitalocean_spaces).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|allowed_methods|`list<item: utf8, nullable>`|
|allowed_origins|`list<item: utf8, nullable>`|
|allowed_headers|`list<item: utf8, nullable>`|
|expose_headers|`list<item: utf8, nullable>`|
|id|`utf8`|
|max_age_seconds|`int64`|