
# Table: k8s_core_resource_quota_scope_selector_match_expressions
A scoped-resource selector requirement is a selector that contains values, a scope name, and an operator that relates the scope name and values.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|resource_quota_cq_id|uuid|Unique CloudQuery ID of k8s_core_resource_quotas table (FK)|
|scope_name|text|The name of the scope that the selector applies to.|
|operator|text|Represents a scope's relationship to a set of values. Valid operators are In, NotIn, Exists, DoesNotExist.|
|values|text[]|An array of string values|
