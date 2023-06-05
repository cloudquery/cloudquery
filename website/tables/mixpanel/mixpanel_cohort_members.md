# Table: mixpanel_cohort_members

This table shows data for Mixpanel Cohort Members.

https://developer.mixpanel.com/reference/engage-query

The composite primary key for this table is (**cohort_id**, **distinct_id**).

## Relations

This table depends on [mixpanel_cohorts](mixpanel_cohorts).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`int64`|
|cohort_id (PK)|`int64`|
|distinct_id (PK)|`utf8`|
|properties|`json`|