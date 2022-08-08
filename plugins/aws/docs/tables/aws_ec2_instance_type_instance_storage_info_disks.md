
# Table: aws_ec2_instance_type_instance_storage_info_disks
Describes a disk.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_type_cq_id|uuid|Unique CloudQuery ID of aws_ec2_instance_types table (FK)|
|count|bigint|The number of disks with this configuration.|
|size_in_gb|bigint|The size of the disk in GB.|
|type|text|The type of disk.|
