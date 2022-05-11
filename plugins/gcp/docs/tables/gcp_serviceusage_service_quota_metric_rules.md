
# Table: gcp_serviceusage_service_quota_metric_rules
Bind API methods to metrics
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_serviceusage_services table (FK)|
|metric_costs|jsonb|Metrics to update when the selected methods are called, and the associated cost applied to each metric|
|selector|text|Selects the methods to which this rule applies|
