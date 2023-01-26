# Table: launchdarkly_project_flag_defaults

https://apidocs.launchdarkly.com/tag/Projects#operation/getFlagDefaultsByProject

The primary key for this table is **project_id**.

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
|tags|StringArray|
|temporary|Bool|
|default_client_side_availability|JSON|
|boolean_defaults|JSON|