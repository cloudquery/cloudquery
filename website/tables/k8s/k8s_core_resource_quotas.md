# Table: k8s_core_resource_quotas

This table shows data for Kubernetes (K8s) Core Resource Quotas.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|utf8|
|_cq_sync_time|timestamp[us, tz=UTC]|
|_cq_id|uuid|
|_cq_parent_id|uuid|
|context|utf8|
|kind|utf8|
|api_version|utf8|
|name|utf8|
|namespace|utf8|
|uid (PK)|utf8|
|resource_version|utf8|
|generation|int64|
|deletion_grace_period_seconds|int64|
|labels|json|
|annotations|json|
|owner_references|json|
|finalizers|list<item: utf8, nullable>|
|spec_hard|json|
|spec_scopes|list<item: utf8, nullable>|
|spec_scope_selector|json|
|status_hard|json|
|status_used|json|