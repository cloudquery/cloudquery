
# Table: aws_firehose_delivery_stream_redshift_destination
Describes a destination in Amazon Redshift
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|delivery_stream_cq_id|uuid|Unique CloudQuery ID of aws_firehose_delivery_streams table (FK)|
|processing_configuration_processors|jsonb|Describes a data processing configuration|
|cluster_j_db_c_url|text|The database connection string|
|copy_command_data_table_name|text|The name of the target table|
|copy_command_copy_options|text|Optional parameters to use with the Amazon Redshift COPY command|
|copy_command_data_table_columns|text|A comma-separated list of column names|
|role_arn|text|The Amazon Resource Name (ARN) of the AWS credentials|
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
|username|text|The name of the user|
|cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging|
|cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|processing_configuration_enabled|boolean|Enables or disables data processing|
|retry_options_duration_in_seconds|bigint|The length of time during which Kinesis Data Firehose retries delivery after a failure, starting from the initial request and including the first attempt|
|s3_backup_bucket_arn|text|The ARN of the S3 bucket|
|s3_backup_buffering_hints_interval_in_seconds|bigint|Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination|
|s3_backup_buffering_hints_size_in_mb_s|bigint|Buffer incoming data to the specified size, in MiBs, before delivering it to the destination|
|s3_backup_compression_format|text|The compression format|
|s3_backup_kms_encryption_config_aws_kms_key_arn|text|The Amazon Resource Name (ARN) of the encryption key|
|s3_backup_no_encryption_config|text|Specifically override existing encryption information to ensure that no encryption is used|
|s3_backup_role_arn|text|The Amazon Resource Name (ARN) of the AWS credentials|
|s3_backup_cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging|
|s3_backup_cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|s3_backup_cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|s3_backup_error_output_prefix|text|A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3|
|s3_backup_prefix|text|The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered Amazon S3 files|
|s3_backup_mode|text|The Amazon S3 backup mode|
