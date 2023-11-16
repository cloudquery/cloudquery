# Table: k8s_core_replication_controllers

This table shows data for Kubernetes (K8s) Core Replication Controllers.

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
|spec_min_ready_seconds|`int64`|
|spec_selector|`json`|
|spec_template|`json`|
|status_replicas|`int64`|
|status_fully_labeled_replicas|`int64`|
|status_ready_replicas|`int64`|
|status_available_replicas|`int64`|
|status_observed_generation|`int64`|
|status_conditions|`json`|