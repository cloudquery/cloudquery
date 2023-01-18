# Table: mixpanel_funnels

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on mixpanel_funnels:
  - [mixpanel_funnel_reports](mixpanel_funnel_reports.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|project_id|Int|
|funnnel_id|Int|
|name|String|