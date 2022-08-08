
# Table: k8s_apps_stateful_set_status_conditions
StatefulSetCondition describes the state of a statefulset at a certain point.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|stateful_set_cq_id|uuid|Unique CloudQuery ID of k8s_apps_stateful_sets table (FK)|
|type|text|Type of statefulset condition.|
|status|text|Status of the condition, one of True, False, Unknown.|
|reason|text|The reason for the condition's last transition.|
|message|text|A human readable message indicating details about the transition.|
