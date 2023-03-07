# Table: k8s_policy_pod_disruption_budgets

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
|spec_min_available|JSON|
|spec_selector|JSON|
|spec_max_unavailable|JSON|
|spec_unhealthy_pod_eviction_policy|String|
|status_observed_generation|Int|
|status_disrupted_pods|JSON|
|status_disruptions_allowed|Int|
|status_current_healthy|Int|
|status_desired_healthy|Int|
|status_expected_pods|Int|
|status_conditions|JSON|