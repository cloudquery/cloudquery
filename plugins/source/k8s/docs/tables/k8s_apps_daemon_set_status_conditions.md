
# Table: k8s_apps_daemon_set_status_conditions
DaemonSetCondition describes the state of a DaemonSet at a certain point.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|daemon_set_cq_id|uuid|Unique CloudQuery ID of k8s_apps_daemon_sets table (FK)|
|type|text|Type of DaemonSet condition.|
|status|text|Status of the condition, one of True, False, Unknown.|
|reason|text|The reason for the condition's last transition.|
|message|text|A human readable message indicating details about the transition.|
