
# Table: aws_ec2_instance_status_events
Any scheduled events associated with the instance.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_status_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instance_statuses table (FK)|
|code|text|The event code.|
|description|text|A description of the event.|
|id|text|The ID of the event.|
|not_after|timestamp without time zone|The latest scheduled end time for the event.|
|not_before|timestamp without time zone|The earliest scheduled start time for the event.|
|not_before_deadline|timestamp without time zone|The deadline for starting the event.|
