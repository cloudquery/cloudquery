# Table: aws_sagemaker_training_jobs



The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|tags|JSON|
|algorithm_specification|JSON|
|creation_time|Timestamp|
|model_artifacts|JSON|
|resource_config|JSON|
|secondary_status|String|
|stopping_condition|JSON|
|training_job_name|String|
|training_job_status|String|
|auto_ml_job_arn|String|
|billable_time_in_seconds|Int|
|checkpoint_config|JSON|
|debug_hook_config|JSON|
|debug_rule_configurations|JSON|
|debug_rule_evaluation_statuses|JSON|
|enable_inter_container_traffic_encryption|Bool|
|enable_managed_spot_training|Bool|
|enable_network_isolation|Bool|
|environment|JSON|
|experiment_config|JSON|
|failure_reason|String|
|final_metric_data_list|JSON|
|hyper_parameters|JSON|
|input_data_config|JSON|
|labeling_job_arn|String|
|last_modified_time|Timestamp|
|output_data_config|JSON|
|profiler_config|JSON|
|profiler_rule_configurations|JSON|
|profiler_rule_evaluation_statuses|JSON|
|profiling_status|String|
|retry_strategy|JSON|
|role_arn|String|
|secondary_status_transitions|JSON|
|tensor_board_output_config|JSON|
|training_end_time|Timestamp|
|training_start_time|Timestamp|
|training_time_in_seconds|Int|
|tuning_job_arn|String|
|vpc_config|JSON|
|warm_pool_status|JSON|
|result_metadata|JSON|