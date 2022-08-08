
# Table: aws_lightsail_disk_snapshot
Describes a block storage disk snapshot
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|disk_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_disks table (FK)|
|arn|text|The Amazon Resource Name (ARN) of the disk snapshot|
|created_at|timestamp without time zone|The date when the disk snapshot was created|
|from_disk_arn|text|The Amazon Resource Name (ARN) of the source disk from which the disk snapshot was created|
|from_disk_name|text|The unique name of the source disk from which the disk snapshot was created|
|from_instance_arn|text|The Amazon Resource Name (ARN) of the source instance from which the disk (system volume) snapshot was created|
|from_instance_name|text|The unique name of the source instance from which the disk (system volume) snapshot was created|
|is_from_auto_snapshot|boolean|A Boolean value indicating whether the snapshot was created from an automatic snapshot|
|location_availability_zone|text|The Availability Zone|
|location_region_name|text|The AWS Region name|
|name|text|The name of the disk snapshot (eg, my-disk-snapshot)|
|progress|text|The progress of the snapshot|
|resource_type|text|The Lightsail resource type (eg, DiskSnapshot)|
|size_in_gb|integer|The size of the disk in GB|
|state|text|The status of the disk snapshot operation|
|support_code|text|The support code|
|tags|jsonb|The tag keys and optional values for the resource|
