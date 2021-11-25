
# Table: aws_sagemaker_training_job_debug_rule_configurations
Configuration information for SageMaker Debugger rules for debugging
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|training_job_cq_id|uuid|Unique CloudQuery ID of aws_sagemaker_training_jobs table (FK)|
|rule_configuration_name|text|The name of the rule configuration|
|rule_evaluator_image|text|The Amazon Elastic Container (ECR) Image for the managed rule evaluation.  This member is required.|
|instance_type|text|The instance type to deploy a Debugger custom rule for debugging a training job.|
|local_path|text|Path to local storage location for output of rules|
|rule_parameters|jsonb|Runtime configuration for rule container.|
|s3_output_path|text|Path to Amazon S3 storage location for rules.|
|volume_size_in_gb|integer|The size, in GB, of the ML storage volume attached to the processing instance.|
