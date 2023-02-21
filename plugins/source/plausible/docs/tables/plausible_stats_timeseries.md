# Table: plausible_stats_timeseries

https://plausible.io/docs/stats-api#get-apiv1statstimeseries

The composite primary key for this table is (**site_id**, **date**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|site_id (PK)|String|
|date (PK)|Timestamp|
|visitors|Int|
|page_views|Int|
|bounce_rate|Int|
|visit_duration|Int|
|visits|Int|