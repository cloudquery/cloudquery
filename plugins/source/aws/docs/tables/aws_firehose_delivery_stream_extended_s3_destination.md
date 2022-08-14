
# Table: aws_firehose_delivery_stream_extended_s3_destination
Describes a destination in Amazon S3
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|delivery_stream_cq_id|uuid|Unique CloudQuery ID of aws_firehose_delivery_streams table (FK)|
|processing_configuration_processors|jsonb|Describes a data processing configuration|
|bucket_arn|text|The ARN of the S3 bucket|
|buffering_hints_interval_in_seconds|bigint|Buffer incoming data for the specified period of time, in seconds, before delivering it to the destination|
|buffering_hints_size_in_mb_s|bigint|Buffer incoming data to the specified size, in MiBs, before delivering it to the destination|
|compression_format|text|The compression format|
|encryption_configuration_kms_encryption_config_aws_kms_key_arn|text|The Amazon Resource Name (ARN) of the encryption key|
|encryption_configuration_no_encryption_config|text|Specifically override existing encryption information to ensure that no encryption is used|
|role_arn|text|The Amazon Resource Name (ARN) of the AWS credentials|
|cloud_watch_logging_options_enabled|boolean|Enables or disables CloudWatch logging|
|cloud_watch_logging_options_log_group_name|text|The CloudWatch group name for logging|
|cloud_watch_logging_options_log_stream_name|text|The CloudWatch log stream name for logging|
|enabled|boolean|Defaults to true|
|deserializer_hive_json_ser_de_timestamp_formats|text[]|Indicates how you want Kinesis Data Firehose to parse the date and timestamps that may be present in your input data JSON|
|deserializer_open_x_json_ser_de_case_insensitive|boolean|When set to true, which is the default, Kinesis Data Firehose converts JSON keys to lowercase before deserializing them|
|deserializer_open_x_json_ser_de_column_to_json_key_mappings|jsonb|Maps column names to JSON keys that aren't identical to the column names|
|deserializer_open_x_json_ser_de_convert_dots_to_underscores|boolean|When set to true, specifies that the names of the keys include dots and that you want Kinesis Data Firehose to replace them with underscores|
|serializer_orc_ser_de_block_size_bytes|bigint|The Hadoop Distributed File System (HDFS) block size|
|serializer_orc_ser_de_bloom_filter_columns|text[]|The column names for which you want Kinesis Data Firehose to create bloom filters|
|serializer_orc_ser_de_bloom_filter_false_positive_probability|float|The Bloom filter false positive probability (FPP)|
|serializer_orc_ser_de_compression|text|The compression code to use over data blocks|
|serializer_orc_ser_de_dictionary_key_threshold|float|Represents the fraction of the total number of non-null rows|
|serializer_orc_ser_de_enable_padding|boolean|Set this to true to indicate that you want stripes to be padded to the HDFS block boundaries|
|serializer_orc_ser_de_format_version|text|The version of the file to write|
|serializer_orc_ser_de_padding_tolerance|float|A number between 0 and 1 that defines the tolerance for block padding as a decimal fraction of stripe size|
|serializer_orc_ser_de_row_index_stride|bigint|The number of rows between index entries|
|serializer_orc_ser_de_stripe_size_bytes|bigint|The number of bytes in each stripe|
|serializer_parquet_ser_de_block_size_bytes|bigint|The Hadoop Distributed File System (HDFS) block size|
|serializer_parquet_ser_de_compression|text|The compression code to use over data blocks|
|serializer_parquet_ser_de_enable_dictionary_compression|boolean|Indicates whether to enable dictionary compression|
|serializer_parquet_ser_de_max_padding_bytes|bigint|The maximum amount of padding to apply|
|serializer_parquet_ser_de_page_size_bytes|bigint|The Parquet page size|
|serializer_parquet_ser_de_writer_version|text|Indicates the version of row format to output|
|schema_configuration_catalog_id|text|The ID of the AWS Glue Data Catalog|
|schema_configuration_database_name|text|Specifies the name of the AWS Glue database that contains the schema for the output data|
|schema_configuration_region|text|If you don't specify an AWS Region, the default is the current Region|
|schema_configuration_role_arn|text|The role that Kinesis Data Firehose can use to access AWS Glue|
|schema_configuration_table_name|text|Specifies the AWS Glue table that contains the column information that constitutes your data schema|
|schema_configuration_version_id|text|Specifies the table version for the output data schema|
|dynamic_partitioning_enabled|boolean|Specifies that the dynamic partitioning is enabled for this Kinesis Data Firehose delivery stream|
|dynamic_partitioning_retry_options_duration_in_seconds|bigint|The period of time during which Kinesis Data Firehose retries to deliver data to the specified Amazon S3 prefix|
|error_output_prefix|text|A prefix that Kinesis Data Firehose evaluates and adds to failed records before writing them to S3|
|prefix|text|The "YYYY/MM/DD/HH" time format prefix is automatically used for delivered Amazon S3 files|
|processing_configuration_enabled|boolean|Enables or disables data processing|
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
