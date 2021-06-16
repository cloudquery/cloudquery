
# Table: aws_cloudtrail_trail_event_selectors
Use event selectors to further specify the management and data event settings for your trail.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|trail_id|uuid|Unique ID of aws_cloudtrail_trails table (FK)|
|exclude_management_event_sources|text[]|An optional list of service event sources from which you do not want management events to be logged on your trail.|
|include_management_events|boolean|Specify if you want your event selector to include management events for your trail.|
|read_write_type|text|Specify if you want your trail to log read-only events, write-only events, or all.|
