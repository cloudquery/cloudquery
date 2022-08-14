
# Table: aws_autoscaling_launch_configuration_block_device_mappings
Describes a block device mapping.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|launch_configuration_cq_id|uuid|Unique CloudQuery ID of aws_autoscaling_launch_configurations table (FK)|
|device_name|text|The device name exposed to the EC2 instance (for example, /dev/sdh or xvdh). For more information, see Device Naming on Linux Instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/device_naming.html) in the Amazon EC2 User Guide for Linux Instances.|
|ebs_delete_on_termination|boolean|Indicates whether the volume is deleted on instance termination. For Amazon EC2 Auto Scaling, the default value is true.|
|ebs_encrypted|boolean|Specifies whether the volume should be encrypted.|
|ebs_iops|integer|The number of I/O operations per second (IOPS) to provision for the volume.|
|ebs_snapshot_id|text|The snapshot ID of the volume to use. You must specify either a VolumeSize or a SnapshotId.|
|ebs_volume_size|integer|The volume size, in Gibibytes (GiB).|
|ebs_volume_type|text|The volume type, which can be standard for Magnetic, io1 for Provisioned IOPS SSD, gp2 for General Purpose SSD, st1 for Throughput Optimized HDD, or sc1 for Cold HDD. For more information, see Amazon EBS Volume Types (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSVolumeTypes.html) in the Amazon EC2 User Guide for Linux Instances. Valid Values: standard | io1 | gp2 | st1 | sc1|
|no_device|boolean|Setting this value to true suppresses the specified device included in the block device mapping of the AMI.|
|virtual_name|text|The name of the virtual device (for example, ephemeral0). You can specify either VirtualName or Ebs, but not both.|
