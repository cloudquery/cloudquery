# Table: mixpanel_funnel_reports

This table shows data for Mixpanel Funnel Reports.

https://developer.mixpanel.com/reference/funnels-query

The composite primary key for this table is (**funnel_id**, **date**).

## Relations

This table depends on [mixpanel_funnels](mixpanel_funnels).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id|int64|
|funnel_id (PK)|int64|
|date (PK)|timestamp[us, tz=UTC]|
|steps|json|
|analysis|json|