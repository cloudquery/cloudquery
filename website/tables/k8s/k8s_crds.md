# Table: k8s_crds

This table shows data for Kubernetes (K8s) Custom Resource Definitions (CRDs).

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
|spec_group|`utf8`|
|spec_names|`json`|
|spec_scope|`utf8`|
|spec_versions|`json`|
|spec_conversion|`json`|
|spec_preserve_unknown_fields|`bool`|
|status_conditions|`json`|
|status_accepted_names|`json`|
|status_stored_versions|`list<item: utf8, nullable>`|