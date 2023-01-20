# Table: mixpanel_cohorts

The primary key for this table is **_cq_id**.

## Relations

The following tables depend on mixpanel_cohorts:
  - [mixpanel_cohort_members](mixpanel_cohort_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id (PK)|UUID|
|_cq_parent_id|UUID|
|count|Int|
|is_visible|Bool|
|description|String|
|created|Timestamp|
|project_id|Int|
|id|Int|
|name|String|