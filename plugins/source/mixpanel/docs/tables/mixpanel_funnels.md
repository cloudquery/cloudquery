# Table: mixpanel_funnels

https://developer.mixpanel.com/reference/funnels-list-saved

The primary key for this table is **funnel_id**.

## Relations

The following tables depend on mixpanel_funnels:
  - [mixpanel_funnel_reports](mixpanel_funnel_reports.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|Int|
|funnel_id (PK)|Int|
|name|String|