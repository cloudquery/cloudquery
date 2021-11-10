
# Table: k8s_apps_daemon_set_selector_match_expressions
A label selector requirement is a selector that contains values, a key, and an operator that relates the key and values.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|daemon_set_cq_id|uuid|Unique CloudQuery ID of k8s_apps_daemon_sets table (FK)|
|key|text|key is the label key that the selector applies to. +patchMergeKey=key +patchStrategy=merge|
|operator|text|operator represents a key's relationship to a set of values. Valid operators are In, NotIn, Exists and DoesNotExist.|
|values|text[]|values is an array of string values|
