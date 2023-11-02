# Table: aws_sagemaker_training_jobs

This table shows data for Amazon SageMaker Training Jobs.

https://docs.aws.amazon.com/sagemaker/latest/APIReference/API_DescribeTrainingJob.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|algorithm_specification|`json`|
|creation_time|`timestamp[us, tz=UTC]`|
|model_artifacts|`json`|
|resource_config|`json`|
|secondary_status|`utf8`|
|stopping_condition|`json`|
|training_job_arn|`utf8`|
|training_job_name|`utf8`|
|training_job_status|`utf8`|
|auto_ml_job_arn|`utf8`|
|billable_time_in_seconds|`int64`|
|checkpoint_config|`json`|
|debug_hook_config|`json`|
|debug_rule_configurations|`json`|
|debug_rule_evaluation_statuses|`json`|
|enable_inter_container_traffic_encryption|`bool`|
|enable_managed_spot_training|`bool`|
|enable_network_isolation|`bool`|
|environment|`json`|
|experiment_config|`json`|
|failure_reason|`utf8`|
|final_metric_data_list|`json`|
|hyper_parameters|`json`|
|input_data_config|`json`|
|labeling_job_arn|`utf8`|
|last_modified_time|`timestamp[us, tz=UTC]`|
|output_data_config|`json`|
|profiler_config|`json`|
|profiler_rule_configurations|`json`|
|profiler_rule_evaluation_statuses|`json`|
|profiling_status|`utf8`|
|retry_strategy|`json`|
|role_arn|`utf8`|
|secondary_status_transitions|`json`|
|tensor_board_output_config|`json`|
|training_end_time|`timestamp[us, tz=UTC]`|
|training_start_time|`timestamp[us, tz=UTC]`|
|training_time_in_seconds|`int64`|
|tuning_job_arn|`utf8`|
|vpc_config|`json`|
|warm_pool_status|`json`|
|result_metadata|`json`|