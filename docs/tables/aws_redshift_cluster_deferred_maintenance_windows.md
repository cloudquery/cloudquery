
# Table: aws_redshift_cluster_deferred_maintenance_windows
Describes a deferred maintenance window .
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid|Unique ID of aws_redshift_clusters table (FK)|
|defer_maintenance_end_time|timestamp without time zone|A timestamp for the end of the time period when we defer maintenance.|
|defer_maintenance_identifier|text|A unique identifier for the maintenance window.|
|defer_maintenance_start_time|timestamp without time zone|A timestamp for the beginning of the time period when we defer maintenance.|
