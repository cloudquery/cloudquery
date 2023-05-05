# Table: digitalocean_regions

This table shows data for DigitalOcean Regions.

https://docs.digitalocean.com/reference/api/api-reference/#tag/Regions

The primary key for this table is **slug**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|slug (PK)|String|
|name|String|
|sizes|StringArray|
|available|Bool|
|features|StringArray|