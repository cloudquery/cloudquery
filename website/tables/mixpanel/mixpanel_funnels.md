# Table: mixpanel_funnels

This table shows data for Mixpanel Funnels.

https://developer.mixpanel.com/reference/funnels-list-saved

The primary key for this table is **funnel_id**.

## Relations

The following tables depend on mixpanel_funnels:
  - [mixpanel_funnel_reports](mixpanel_funnel_reports)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|int64|
|funnel_id (PK)|int64|
|name|utf8|