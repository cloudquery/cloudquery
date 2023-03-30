# Table: k8s_core_config_maps

This table shows data for Kubernetes (K8s) Core Config Maps.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|context|String|
|kind|String|
|api_version|String|
|name|String|
|namespace|String|
|uid (PK)|String|
|resource_version|String|
|generation|Int|
|deletion_grace_period_seconds|Int|
|labels|JSON|
|annotations|JSON|
|owner_references|JSON|
|finalizers|StringArray|
|immutable|Bool|
|data|JSON|
|binary_data|JSON|