# Table: digitalocean_space_cors



The primary key for this table is **id**.

## Relations
This table depends on [digitalocean_spaces](digitalocean_spaces.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|id (PK)|String|
|allowed_methods|StringArray|
|allowed_origins|StringArray|
|allowed_headers|StringArray|
|expose_headers|StringArray|
|max_age_seconds|Int|