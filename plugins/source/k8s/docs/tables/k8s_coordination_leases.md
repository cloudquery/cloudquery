# Table: k8s_coordination_leases

This table shows data for Kubernetes (K8s) Coordination Leases.

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
|spec_holder_identity|`utf8`|
|spec_lease_duration_seconds|`int64`|
|spec_acquire_time|`json`|
|spec_renew_time|`json`|
|spec_lease_transitions|`int64`|