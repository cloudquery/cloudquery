
# Table: aws_ec2_image_block_device_mappings
Describes a block device mapping.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|image_id|uuid|Unique ID of aws_ec2_images table (FK)|
|device_name|text|The device name (for example, /dev/sdh or xvdh).|
|ebs_delete_on_termination|boolean|Indicates whether the EBS volume is deleted on instance termination.|
|ebs_encrypted|boolean|Indicates whether the encryption state of an EBS volume is changed while being restored from a backing snapshot.|
|ebs_iops|integer|The number of I/O operations per second (IOPS).|
|ebs_kms_key_id|text|Identifier (key ID, key alias, ID ARN, or alias ARN) for a customer managed CMK under which the EBS volume is encrypted.|
|ebs_outpost_arn|text|The ARN of the Outpost on which the snapshot is stored.|
|ebs_snapshot_id|text|The ID of the snapshot.|
|ebs_throughput|integer|The throughput that the volume supports, in MiB/s.|
|ebs_volume_size|integer|The size of the volume, in GiBs.|
|ebs_volume_type|text|The volume type.|
|no_device|text|To omit the device from the block device mapping, specify an empty string.|
|virtual_name|text|The virtual device name (ephemeralN).|
