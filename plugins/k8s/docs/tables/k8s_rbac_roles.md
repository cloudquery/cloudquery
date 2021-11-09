
# Table: k8s_rbac_roles
Role is a namespaced, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|context|text|Name of the context from k8s configuration.|
|kind|text||
|api_version|text||
|name|text||
|generate_name|text||
|namespace|text||
|self_link|text||
|uid|text||
|resource_version|text||
|generation|bigint||
|deletion_grace_period_seconds|bigint||
|labels|jsonb||
|annotations|jsonb||
|owner_references|jsonb||
|finalizers|text[]||
|cluster_name|text||
|managed_fields|jsonb||
