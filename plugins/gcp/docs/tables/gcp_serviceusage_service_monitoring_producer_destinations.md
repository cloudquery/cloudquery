
# Table: gcp_serviceusage_service_monitoring_producer_destinations
Configuration of a specific monitoring destination (the producer project or the consumer project)
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_serviceusage_services table (FK)|
|metrics|text[]|Types of the metrics to report to this monitoring destination|
|monitored_resource|text|The monitored resource type|
