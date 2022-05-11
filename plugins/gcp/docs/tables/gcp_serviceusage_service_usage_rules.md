
# Table: gcp_serviceusage_service_usage_rules
Usage configuration rules for the service
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_serviceusage_services table (FK)|
|allow_unregistered_calls|boolean|If true, the selected method allows unregistered calls, eg|
|selector|text|Selects the methods to which this rule applies|
|skip_service_control|boolean|If true, the selected method should skip service control and the control plane features, such as quota and billing, will not be available|
