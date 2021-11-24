
# Table: aws_ec2_ebs_snapshots
Describes a snapshot.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|create_volume_permissions|jsonb||
|data_encryption_key_id|text|The data encryption key identifier for the snapshot|
|description|text|The description for the snapshot.|
|encrypted|boolean|Indicates whether the snapshot is encrypted.|
|kms_key_id|text|The Amazon Resource Name (ARN) of the AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the volume encryption key for the parent volume.|
|outpost_arn|text|The ARN of the AWS Outpost on which the snapshot is stored|
|owner_alias|text|The AWS owner alias, from an Amazon-maintained list (amazon)|
|owner_id|text|The AWS account ID of the EBS snapshot owner.|
|progress|text|The progress of the snapshot, as a percentage.|
|snapshot_id|text|The ID of the snapshot|
|start_time|timestamp without time zone|The time stamp when the snapshot was initiated.|
|state|text|The snapshot state.|
|state_message|text|Encrypted Amazon EBS snapshots are copied asynchronously|
|tags|jsonb|Any tags assigned to the snapshot.|
|volume_id|text|The ID of the volume that was used to create the snapshot|
|volume_size|integer|The size of the volume, in GiB.|
