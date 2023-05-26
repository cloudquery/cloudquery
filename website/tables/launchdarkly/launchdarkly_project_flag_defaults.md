# Table: launchdarkly_project_flag_defaults

This table shows data for LaunchDarkly Project Flag Defaults.

https://apidocs.launchdarkly.com/tag/Projects#operation/getFlagDefaultsByProject

The primary key for this table is **project_id**.

## Relations

This table depends on [launchdarkly_projects](launchdarkly_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|project_id (PK)|utf8|
|tags|list<item: utf8, nullable>|
|temporary|bool|
|default_client_side_availability|json|
|boolean_defaults|json|