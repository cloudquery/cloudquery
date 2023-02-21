# Table: mixpanel_funnel_reports

https://developer.mixpanel.com/reference/funnels-query

The composite primary key for this table is (**funnel_id**, **date**).

## Relations

This table depends on [mixpanel_funnels](mixpanel_funnels.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|Int|
|funnel_id (PK)|Int|
|date (PK)|Timestamp|
|steps|JSON|
|analysis|JSON|