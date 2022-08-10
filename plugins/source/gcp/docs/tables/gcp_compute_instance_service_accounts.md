
# Table: gcp_compute_instance_service_accounts
A service account
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique ID of gcp_compute_instances table (FK)|
|instance_id|text||
|email|text|Email address of the service account|
|scopes|text[]|The list of scopes to be made available for this service account|
