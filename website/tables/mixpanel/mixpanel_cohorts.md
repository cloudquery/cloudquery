# Table: mixpanel_cohorts

This table shows data for Mixpanel Cohorts.

https://developer.mixpanel.com/reference/cohorts-list

The primary key for this table is **id**.

## Relations

The following tables depend on mixpanel_cohorts:
  - [mixpanel_cohort_members](mixpanel_cohort_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|count|int64|
|is_visible|int64|
|description|utf8|
|created|timestamp[us, tz=UTC]|
|project_id|int64|
|id (PK)|int64|
|name|utf8|