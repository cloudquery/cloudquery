
# Table: aws_sagemaker_training_job_input_data_config
A channel is a named input source that training algorithms can consume.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|training_job_cq_id|uuid|Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)|
|channel_name|text|The name of the channel. |
|data_source_file_directory_path|text|The full path to the directory to associate with the channel. |
|data_source_file_system_access_mode|text|The access mode of the mount of the directory associated with the channel|
|data_source_file_system_id|text|The file system id. |
|data_source_file_system_type|text|The file system type. |
|data_source_s3_data_type|text|If you choose S3Prefix, S3Uri identifies a key name prefix|
|data_source_s3_uri|text|Depending on the value specified for the S3DataType, identifies either a key name prefix or a manifest|
|data_source_attribute_names|text[]|A list of one or more attribute names to use that are found in a specified augmented manifest file.|
|data_source_s3_data_distribution_type|text|If you want Amazon SageMaker to replicate the entire dataset on each ML compute instance that is launched for model training, specify FullyReplicated|
|compression_type|text|If training data is compressed, the compression type|
|content_type|text|The MIME type of the data.|
|input_mode|text|(Optional) The input mode to use for the data channel in a training job|
|record_wrapper_type|text|Specify RecordIO as the value when input data is in raw format but the training algorithm requires the RecordIO format|
|shuffle_config_seed|bigint|Determines the shuffling order in ShuffleConfig value. |
