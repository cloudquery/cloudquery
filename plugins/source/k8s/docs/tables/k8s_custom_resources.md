# Table: k8s_custom_resources

This table shows data for Kubernetes (K8s) Custom Resources.

The primary key for this table is **uid**.

## Relations

This table depends on [k8s_crds](k8s_crds.md).

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
|labels|`json`|
|annotations|`json`|
|owner_references|`json`|
|finalizers|`list<item: utf8, nullable>`|
|spec|`json`|
|status|`json`|