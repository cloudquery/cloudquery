
# Table: gcp_cloudrun_service_status_conditions
Condition defines a generic condition for a Resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_services table (FK)|
|last_transition_time|text|Optional|
|message|text|Optional|
|reason|text|Optional|
|severity|text|Optional|
|status|text|Status of the condition, one of True, False, Unknown|
|type|text|type is used to communicate the status of the reconciliation process|
