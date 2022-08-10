
# Table: gcp_cloudrun_service_status_traffic
TrafficTarget holds a single entry of the routing table for a Route
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|service_cq_id|uuid|Unique CloudQuery ID of gcp_cloudrun_services table (FK)|
|configuration_name|text|ConfigurationName of a configuration to whose latest revision we will send this portion of traffic|
|latest_revision|boolean|Optional|
|percent|bigint|Percent specifies percent of the traffic to this Revision or Configuration|
|revision_name|text|RevisionName of a specific revision to which to send this portion of traffic|
|tag|text|Optional|
|url|text|Output only|
