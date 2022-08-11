
# Table: aws_sagemaker_training_job_profiler_rule_evaluation_statuses
Information about the status of the rule evaluation.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|training_job_cq_id|uuid|Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)|
|last_modified_time|timestamp without time zone|Timestamp when the rule evaluation status was last modified.|
|rule_configuration_name|text|The name of the rule configuration.|
|rule_evaluation_job_arn|text|The Amazon Resource Name (ARN) of the rule evaluation job.|
|rule_evaluation_status|text|Status of the rule evaluation.|
|status_details|text|Details from the rule evaluation.|
