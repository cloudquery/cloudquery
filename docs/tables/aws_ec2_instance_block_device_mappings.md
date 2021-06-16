
# Table: aws_ec2_instance_block_device_mappings
Describes a block device mapping.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|instance_id|uuid|Unique ID of aws_ec2_instances table (FK)|
|device_name|text|The device name (for example, /dev/sdh or xvdh).|
|ebs_attach_time|timestamp without time zone|The time stamp when the attachment initiated.|
|ebs_delete_on_termination|boolean|Indicates whether the volume is deleted on instance termination.|
|ebs_status|text|The attachment state.|
|ebs_volume_id|text|The ID of the EBS volume.|
