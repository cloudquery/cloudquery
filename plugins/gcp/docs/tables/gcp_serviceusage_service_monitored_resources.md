
# Table: gcp_serviceusage_service_monitored_resources
An object that describes the schema of a MonitoredResource object using a type name and a set of labels
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_serviceusage_services table (FK)|
|description|text|Optional|
|display_name|text|Optional|
|labels|jsonb|Required|
|launch_stage|text|Optional|
|name|text|Optional|
|type|text|Required|
