
# Table: k8s_apps_replica_set_status_conditions
ReplicaSetCondition describes the state of a replica set at a certain point.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|replica_set_cq_id|uuid|Unique CloudQuery ID of k8s_apps_replica_sets table (FK)|
|type|text|Type of replica set condition.|
|status|text|Status of the condition, one of True, False, Unknown.|
|reason|text|The reason for the condition's last transition.|
|message|text|A human readable message indicating details about the transition.|
