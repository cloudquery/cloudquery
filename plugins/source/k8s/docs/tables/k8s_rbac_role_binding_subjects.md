
# Table: k8s_rbac_role_binding_subjects
Subject contains a reference to the object or user identities a role binding applies to
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|role_binding_cq_id|uuid|Unique CloudQuery ID of k8s_rbac_role_bindings table (FK)|
|kind|text|Kind of object being referenced|
|api_group|text|APIGroup holds the API group of the referenced subject. Defaults to "" for ServiceAccount subjects. Defaults to "rbac.authorization.k8s.io" for User and Group subjects. +optional|
|name|text|Name of the object being referenced.|
|namespace|text|Namespace of the referenced object|
