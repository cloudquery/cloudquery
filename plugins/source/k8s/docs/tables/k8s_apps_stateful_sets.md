# Table: k8s_apps_stateful_sets

This table shows data for Kubernetes (K8s) Apps Stateful Sets.

The primary key for this table is **uid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|context|`utf8`|
|kind|`utf8`|
|api_version|`utf8`|
|name|`utf8`|
|namespace|`utf8`|
|uid (PK)|`utf8`|
|resource_version|`utf8`|
|generation|`int64`|
|deletion_grace_period_seconds|`int64`|
|labels|`json`|
|annotations|`json`|
|owner_references|`json`|
|finalizers|`list<item: utf8, nullable>`|
|spec_replicas|`int64`|
|spec_selector|`json`|
|spec_template|`json`|
|spec_volume_claim_templates|`json`|
|spec_service_name|`utf8`|
|spec_pod_management_policy|`utf8`|
|spec_update_strategy|`json`|
|spec_revision_history_limit|`int64`|
|spec_min_ready_seconds|`int64`|
|spec_persistent_volume_claim_retention_policy|`json`|
|spec_ordinals|`json`|
|status_observed_generation|`int64`|
|status_replicas|`int64`|
|status_ready_replicas|`int64`|
|status_current_replicas|`int64`|
|status_updated_replicas|`int64`|
|status_current_revision|`utf8`|
|status_update_revision|`utf8`|
|status_collision_count|`int64`|
|status_conditions|`json`|
|status_available_replicas|`int64`|