
# Table: k8s_rbac_role_bindings
RoleBinding references a role, but does not contain it
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|k8s_config_context|text|Name of the context from k8s configuration.|
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
|role_ref_api_group|text|APIGroup is the group for the resource being referenced|
|role_ref_kind|text|Kind is the type of resource being referenced|
|role_ref_name|text|Name is the name of resource being referenced|
