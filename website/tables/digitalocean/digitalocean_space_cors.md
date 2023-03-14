# Table: digitalocean_space_cors

This table shows data for Digitalocean Space CORS.

The primary key for this table is **_cq_id**.

## Relations

This table depends on [digitalocean_spaces](digitalocean_spaces).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|allowed_methods|StringArray|
|allowed_origins|StringArray|
|allowed_headers|StringArray|
|expose_headers|StringArray|
|id|String|
|max_age_seconds|Int|