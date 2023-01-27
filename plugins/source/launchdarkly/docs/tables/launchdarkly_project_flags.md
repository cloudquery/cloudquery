# Table: launchdarkly_project_flags

https://apidocs.launchdarkly.com/tag/Feature-flags#operation/getFeatureFlags

The composite primary key for this table is (**project_id**, **key**).

## Relations

This table depends on [launchdarkly_projects](launchdarkly_projects.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|project_id (PK)|String|
|name|String|
|kind|String|
|description|String|
|key (PK)|String|
|version|Int|
|creation_date|Int|
|include_in_snippet|Bool|
|client_side_availability|JSON|
|variations|JSON|
|temporary|Bool|
|tags|StringArray|
|maintainer_id|String|
|maintainer|JSON|
|goal_ids|StringArray|
|experiments|JSON|
|custom_properties|JSON|
|archived|Bool|
|archived_date|Int|
|defaults|JSON|
|environments|JSON|