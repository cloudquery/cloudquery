# Table: mixpanel_cohorts

https://developer.mixpanel.com/reference/cohorts-list

The primary key for this table is **id**.

## Relations

The following tables depend on mixpanel_cohorts:
  - [mixpanel_cohort_members](mixpanel_cohort_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|count|Int|
|is_visible|Int|
|description|String|
|created|Timestamp|
|project_id|Int|
|id (PK)|Int|
|name|String|