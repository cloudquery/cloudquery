# Table: launchdarkly_project_flags

This table shows data for LaunchDarkly Project Flags.

https://apidocs.launchdarkly.com/tag/Feature-flags#operation/getFeatureFlags

The composite primary key for this table is (**project_id**, **key**).

## Relations

This table depends on [launchdarkly_projects](launchdarkly_projects).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|project_id (PK)|`utf8`|
|name|`utf8`|
|kind|`utf8`|
|description|`utf8`|
|key (PK)|`utf8`|
|version|`int64`|
|creation_date|`int64`|
|include_in_snippet|`bool`|
|client_side_availability|`json`|
|variations|`json`|
|temporary|`bool`|
|tags|`list<item: utf8, nullable>`|
|maintainer_id|`utf8`|
|maintainer|`json`|
|goal_ids|`list<item: utf8, nullable>`|
|experiments|`json`|
|custom_properties|`json`|
|archived|`bool`|
|archived_date|`int64`|
|defaults|`json`|
|environments|`json`|