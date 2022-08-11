
# Table: aws_sagemaker_training_jobs
Provides summary information about a training job.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|auto_ml_job_arn|text|The Amazon Resource Name (ARN) of an AutoML job.|
|billable_time_in_seconds|integer|The billable time in seconds. Billable time refers to the absolute wall-clock time.|
|enable_managed_spot_training|boolean|A Boolean indicating whether managed spot training is enabled (True) or not (False).|
|enable_network_isolation|boolean|If you want to allow inbound or outbound network calls, except for calls between peers within a training cluster for distributed training, choose True. If you enable network isolation for training jobs that are configured to use a VPC, Amazon SageMaker downloads and uploads customer data and model artifacts through the specified VPC, but the training container does not have network access.|
|enable_inter_container_traffic_encryption|boolean|To encrypt all communications between ML compute instances in distributed training, choose True. Encryption provides greater security for distributed training, but training might take longer. How long it takes depends on the amount of communication between compute instances, especially if you use a deep learning algorithms in distributed training.|
|failure_reason|text|If the training job failed, the reason it failed.|
|labeling_job_arn|text|The Amazon Resource Name (ARN) of the Amazon SageMaker Ground Truth labeling job that created the transform or training job.|
|last_modified_time|timestamp without time zone|A timestamp that indicates when the status of the training job was last modified.|
|profiling_status|text|Profiling status of a training job.|
|role_arn|text|The Amazon Web Services Identity and Access Management (IAM) role configured for the training job.|
|secondary_status|text|Provides detailed information about the state of the training job.|
|training_end_time|timestamp without time zone|Indicates the time when the training job ends on training instances.|
|training_start_time|timestamp without time zone|Indicates the time when the training job starts on training instances.|
|training_time_in_seconds|integer|The training time in seconds.|
|tuning_job_arn|text|The Amazon Resource Name (ARN) of the associated hyperparameter tuning job if the training job was launched by a hyperparameter tuning job.|
|checkpoint_config|jsonb|Contains information about the output location for managed spot training checkpoint data.|
|environment|jsonb|The environment variables to set in the Docker container.|
|experiment_config|jsonb|Associates a SageMaker job as a trial component with an experiment and trial.|
|hyper_parameters|jsonb|Algorithm-specific parameters.|
|model_artifacts|jsonb|Information about the Amazon S3 location that is configured for storing model artifacts.|
|output_data_config|jsonb|The S3 path where model artifacts that you configured when creating the job are stored.|
|profiler_config|jsonb|Configuration information for Debugger system monitoring, framework profiling, and storage paths.|
|resource_config|jsonb|Resources, including ML compute instances and ML storage volumes, that are configured for model training.|
|stopping_condition|jsonb|Specifies a limit to how long a model training job can run.|
|tensor_board_output_config|jsonb|Configuration of storage locations for the Debugger TensorBoard output data.|
|vpc_config|jsonb|A VpcConfig object that specifies the VPC that this training job has access to.|
|tags|jsonb|The tags associated with the model.|
|creation_time|timestamp without time zone|A timestamp that shows when the training job was created.|
|arn|text|The Amazon Resource Name (ARN) of the training job.|
|name|text|The name of the training job.|
|training_job_status|text|The status of the training job.|
|secondary_status_transitions|jsonb||
|final_metric_data_list|jsonb||
