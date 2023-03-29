# Table: k8s_crds

This table shows data for Kubernetes (K8s) Custom Resource Definitions (CRDs).

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
|spec_group|String|
|spec_names|JSON|
|spec_scope|String|
|spec_versions|JSON|
|spec_conversion|JSON|
|spec_preserve_unknown_fields|Bool|
|status_conditions|JSON|
|status_accepted_names|JSON|
|status_stored_versions|StringArray|