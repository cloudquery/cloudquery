
# Table: gcp_compute_autoscaler_custom_metric_utilizations
Custom utilization metric policy
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|autoscaler_cq_id|uuid|Unique ID of gcp_compute_autoscalers table (FK)|
|autoscaler_id|text||
|filter|text|A filter string, compatible with a Stackdriver Monitoring filter string|
|metric|text|The identifier (type) of the Stackdriver Monitoring metric The metric cannot have negative values  The metric must have a value type of INT64 or DOUBLE|
|single_instance_assignment|float|per-group metric value that represents the total amount of work to be done or resource usage|
|utilization_target|float|The target value of the metric that autoscaler maintains This must be a positive value A utilization metric scales number of virtual machines handling requests to increase or decrease proportionally to the metric  For example, a good metric to use as a utilization_target is https://wwwgoogleapiscom/compute/v1/instance/network/received_bytes_count The autoscaler works to keep this value constant for each of the instances|
|utilization_target_type|text|Defines how target utilization value is expressed for a Stackdriver Monitoring metric Either GAUGE, DELTA_PER_SECOND, or DELTA_PER_MINUTE|
