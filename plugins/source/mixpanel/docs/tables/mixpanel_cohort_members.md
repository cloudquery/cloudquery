# Table: mixpanel_cohort_members

https://developer.mixpanel.com/reference/engage-query

The composite primary key for this table is (**cohort_id**, **distinct_id**).

## Relations

This table depends on [mixpanel_cohorts](mixpanel_cohorts.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id|Int|
|cohort_id (PK)|Int|
|distinct_id (PK)|String|
|properties|JSON|