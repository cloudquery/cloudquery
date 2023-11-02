# Table: k8s_policy_pod_disruption_budgets

This table shows data for Kubernetes (K8s) Policy Pod Disruption Budgets.

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
|spec_min_available|`json`|
|spec_selector|`json`|
|spec_max_unavailable|`json`|
|spec_unhealthy_pod_eviction_policy|`utf8`|
|status_observed_generation|`int64`|
|status_disrupted_pods|`json`|
|status_disruptions_allowed|`int64`|
|status_current_healthy|`int64`|
|status_desired_healthy|`int64`|
|status_expected_pods|`int64`|
|status_conditions|`json`|