# Table: k8s_apps_deployments

This table shows data for Kubernetes (K8s) Apps Deployments.

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
|spec_replicas|int64|
|spec_selector|json|
|spec_template|json|
|spec_strategy|json|
|spec_min_ready_seconds|int64|
|spec_revision_history_limit|int64|
|spec_paused|bool|
|spec_progress_deadline_seconds|int64|
|status_observed_generation|int64|
|status_replicas|int64|
|status_updated_replicas|int64|
|status_ready_replicas|int64|
|status_available_replicas|int64|
|status_unavailable_replicas|int64|
|status_conditions|json|
|status_collision_count|int64|