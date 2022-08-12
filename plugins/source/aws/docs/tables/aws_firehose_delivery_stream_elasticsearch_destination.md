
# Table: aws_firehose_delivery_stream_elasticsearch_destination
The destination description in Amazon ES
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|delivery_stream_cq_id|uuid|Unique CloudQuery ID of aws_firehose_delivery_streams table (FK)|
|processing_configuration_processors|jsonb|Describes a data processing configuration|
|buffering_hints_interval_in_seconds|bigint|Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination|
|buffering_hints_size_in_mb_s|bigint|Buffer incoming data to the specified size, in MBs, before delivering it to the destination|
|cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging|
|cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|cluster_endpoint|text|The endpoint to use when communicating with the cluster|
|domain_arn|text|The ARN of the Amazon ES domain|
|index_name|text|The Elasticsearch index name|
|index_rotation_period|text|The Elasticsearch index rotation period|
|processing_configuration_enabled|boolean|Enables or disables data processing|
|retry_options_duration_in_seconds|bigint|After an initial failure to deliver to Amazon ES, the total amount of time during which Kinesis Data Firehose retries delivery (including the first attempt)|
|role_arn|text|The Amazon Resource Name (ARN) of the AWS credentials|
|s3_backup_mode|text|The Amazon S3 backup mode|
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
|type_name|text|The Elasticsearch type name|
|vpc_configuration_description_role_arn|text|The ARN of the IAM role that the delivery stream uses to create endpoints in the destination VPC|
|vpc_configuration_description_security_group_ids|text[]|The IDs of the security groups that Kinesis Data Firehose uses when it creates ENIs in the VPC of the Amazon ES destination|
|vpc_configuration_description_subnet_ids|text[]|The IDs of the subnets that Kinesis Data Firehose uses to create ENIs in the VPC of the Amazon ES destination|
|vpc_configuration_description_vpc_id|text|The ID of the Amazon ES destination's VPC|
