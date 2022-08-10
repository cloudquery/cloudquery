
# Table: aws_lightsail_instance_add_ons
Describes an add-on that is enabled for an Amazon Lightsail resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_instances table (FK)|
|name|text|The name of the add-on|
|next_snapshot_time_of_day|text|The next daily time an automatic snapshot will be created|
|snapshot_time_of_day|text|The daily time when an automatic snapshot is created|
|status|text|The status of the add-on|
