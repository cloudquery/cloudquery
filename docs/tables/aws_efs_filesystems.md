
# Table: aws_efs_filesystems
A description of the file system.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|creation_time|timestamp without time zone|The time that the file system was created, in seconds (since 1970-01-01T00:00:00Z).|
|creation_token|text|The opaque string specified in the request.|
|id|text|The ID of the file system, assigned by Amazon EFS.|
|life_cycle_state|text|The lifecycle phase of the file system.|
|number_of_mount_targets|integer|The current number of mount targets that the file system has.|
|owner_id|text|The AWS account that created the file system.|
|performance_mode|text|The performance mode of the file system.|
|size_in_bytes_value|bigint|The latest known metered size (in bytes) of data stored in the file system.|
|size_in_bytes_timestamp|timestamp without time zone|The time at which the size of data, returned in the Value field, was determined.|
|size_in_bytes_value_in_ia|bigint|The latest known metered size (in bytes) of data stored in the Infrequent Access storage class.|
|size_in_bytes_value_in_standard|bigint|The latest known metered size (in bytes) of data stored in the Standard storage class.|
|tags|jsonb|The tags associated with the file system, presented as an array of Tag objects.|
|availability_zone_id|text|The unique and consistent identifier of the Availability Zone in which the file system's One Zone storage classes exist.|
|availability_zone_name|text|Describes the AWS Availability Zone in which the file system is located, and is valid only for file systems using One Zone storage classes.|
|encrypted|boolean|A Boolean value that, if true, indicates that the file system is encrypted.|
|arn|text|The Amazon Resource Name (ARN) for the EFS file system, in the format arn:aws:elasticfilesystem:region:account-id:file-system/file-system-id .|
|kms_key_id|text|The ID of an AWS Key Management Service (AWS KMS) customer master key (CMK) that was used to protect the encrypted file system.|
|name|text|You can add tags to a file system, including a Name tag.|
|provisioned_throughput_in_mibps|float|The amount of provisioned throughput, measured in MiB/s, for the file system.|
|throughput_mode|text|Displays the file system's throughput mode.|
