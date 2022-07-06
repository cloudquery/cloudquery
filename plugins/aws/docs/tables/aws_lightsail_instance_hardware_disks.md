
# Table: aws_lightsail_instance_hardware_disks
Describes a block storage disk.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_instances table (FK)|
|arn|text|The Amazon Resource Name (ARN) of the disk.|
|attached_to|text|The resources to which the disk is attached.|
|created_at|timestamp without time zone|The date when the disk was created.|
|iops|integer|The input/output operations per second (IOPS) of the disk.|
|is_attached|boolean|A Boolean value indicating whether the disk is attached.|
|is_system_disk|boolean|A Boolean value indicating whether this disk is a system disk (has an operating system loaded on it).|
|location_availability_zone|text|The Availability Zone|
|location_region_name|text|The AWS Region name.|
|name|text|The unique name of the disk.|
|path|text|The disk path.|
|resource_type|text|The Lightsail resource type (e.g., Disk).|
|size_in_gb|integer|The size of the disk in GB.|
|state|text|Describes the status of the disk.|
|support_code|text|The support code|
|tags|jsonb|The tag keys and optional values for the resource|
