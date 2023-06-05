# Table: plausible_stats_timeseries

This table shows data for Plausible Stats Timeseries.

https://plausible.io/docs/stats-api#get-apiv1statstimeseries

The composite primary key for this table is (**site_id**, **date**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|site_id (PK)|`utf8`|
|date (PK)|`timestamp[us, tz=UTC]`|
|visitors|`int64`|
|page_views|`int64`|
|bounce_rate|`int64`|
|visit_duration|`int64`|
|visits|`int64`|