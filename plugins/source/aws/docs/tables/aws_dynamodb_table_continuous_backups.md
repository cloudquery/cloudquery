
# Table: aws_dynamodb_table_continuous_backups
Represents the continuous backups and point in time recovery settings on the table.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|table_cq_id|uuid|Unique CloudQuery ID of aws_dynamodb_tables table (FK)|
|continuous_backups_status|text|ContinuousBackupsStatus can be one of the following states: ENABLED, DISABLED |
|earliest_restorable_date_time|timestamp without time zone|Specifies the earliest point in time you can restore your table to|
|latest_restorable_date_time|timestamp without time zone|LatestRestorableDateTime is typically 5 minutes before the current time.|
|point_in_time_recovery_status|text|The current state of point in time recovery:  * ENABLING - Point in time recovery is being enabled.  * ENABLED - Point in time recovery is enabled.  * DISABLED - Point in time recovery is disabled.|
