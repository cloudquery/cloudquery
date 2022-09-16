
# Table: aws_lightsail_instance_snapshot_from_attached_disk_add_ons
Describes an add-on that is enabled for an Amazon Lightsail resource
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_snapshot_from_attached_disk_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_instance_snapshot_from_attached_disks table (FK)|
|name|text|The name of the add-on|
|next_snapshot_time_of_day|text|The next daily time an automatic snapshot will be created|
|snapshot_time_of_day|text|The daily time when an automatic snapshot is created|
|status|text|The status of the add-on|
