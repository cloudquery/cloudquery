# Table: k8s_autoscaling_hpas

This table shows data for Kubernetes (K8s) Autoscaling Horizontal Pod Autoscalers (HPAs).

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
|spec_scale_target_ref|`json`|
|spec_min_replicas|`int64`|
|spec_max_replicas|`int64`|
|spec_target_cpu_utilization_percentage|`int64`|
|status_observed_generation|`int64`|
|status_last_scale_time|`timestamp[us, tz=UTC]`|
|status_current_replicas|`int64`|
|status_desired_replicas|`int64`|
|status_current_cpu_utilization_percentage|`int64`|