
# Table: k8s_apps_deployment_status_conditions
DeploymentCondition describes the state of a deployment at a certain point.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|deployment_cq_id|uuid|Unique CloudQuery ID of k8s_apps_deployments table (FK)|
|type|text|Type of deployment condition.|
|status|text|Status of the condition, one of True, False, Unknown.|
|reason|text|The reason for the condition's last transition.|
|message|text|A human readable message indicating details about the transition.|
