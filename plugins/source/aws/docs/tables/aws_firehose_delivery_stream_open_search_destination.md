
# Table: aws_firehose_delivery_stream_open_search_destination

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|delivery_stream_cq_id|uuid|Unique CloudQuery ID of aws_firehose_delivery_streams table (FK)|
|processing_configuration_processors|jsonb|Describes a data processing configuration|
|buffering_hints_interval_in_seconds|bigint||
|buffering_hints_size_in_mb_s|bigint||
|cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging|
|cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|cluster_endpoint|text||
|domain_arn|text||
|index_name|text||
|index_rotation_period|text||
|processing_configuration_enabled|boolean|Enables or disables data processing|
|retry_options_duration_in_seconds|bigint||
|role_arn|text||
|s3_backup_mode|text||
|s3_destination_bucket_arn|text|The ARN of the S3 bucket|
|s3_destination_buffering_hints_interval_in_seconds|bigint|Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination|
|s3_destination_buffering_hints_size_in_mb_s|bigint|Buffer incoming data to the specified size, in MiBs, before delivering it to the destination|
|s3_destination_compression_format|text|The compression format|
|s3_destination_kms_encryption_config_aws_kms_key_arn|text|The Amazon Resource Name (ARN) of the encryption key|
|s3_destination_no_encryption_config|text|Specifically override existing encryption information to ensure that no encryption is used|
|s3_destination_role_arn|text|The Amazon Resource Name (ARN) of the AWS credentials|
|s3_destination_cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging|
|s3_destination_cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|s3_destination_cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|s3_destination_error_output_prefix|text|A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3|
|s3_destination_prefix|text|The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered Amazon S3 files|
|type_name|text||
|vpc_configuration_description_role_arn|text|The ARN of the IAM role that the delivery stream uses to create endpoints in the destination VPC|
|vpc_configuration_description_security_group_ids|text[]|The IDs of the security groups that Kinesis Data Firehose uses when it creates ENIs in the VPC of the Amazon ES destination|
|vpc_configuration_description_subnet_ids|text[]|The IDs of the subnets that Kinesis Data Firehose uses to create ENIs in the VPC of the Amazon ES destination|
|vpc_configuration_description_vpc_id|text|The ID of the Amazon ES destination's VPC|
