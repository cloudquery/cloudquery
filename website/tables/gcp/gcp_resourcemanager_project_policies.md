# Table: gcp_resourcemanager_project_policies

This table shows data for GCP Resourcemanager Project Policies.

https://cloud.google.com/resource-manager/reference/rest/Shared.Types/Policy

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|project_id|`utf8`|
|audit_configs|`json`|
|bindings|`json`|
|etag|`utf8`|
|version|`int64`|