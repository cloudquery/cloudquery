# Table: k8s_autoscaling_hpas



The primary key for this table is **uid**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|context|String|
|uid (PK)|String|
|kind|String|
|api_version|String|
|name|String|
|namespace|String|
|resource_version|String|
|generation|Int|
|deletion_grace_period_seconds|Int|
|labels|JSON|
|annotations|JSON|
|owner_references|JSON|
|finalizers|StringArray|
|spec_scale_target_ref|JSON|
|spec_min_replicas|Int|
|spec_max_replicas|Int|
|spec_target_cpu_utilization_percentage|Int|
|status_observed_generation|Int|
|status_last_scale_time|Timestamp|
|status_current_replicas|Int|
|status_desired_replicas|Int|
|status_current_cpu_utilization_percentage|Int|