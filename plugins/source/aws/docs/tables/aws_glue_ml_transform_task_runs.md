
# Table: aws_glue_ml_transform_task_runs
The sampling parameters that are associated with the machine learning transform
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|ml_transform_cq_id|uuid|Unique CloudQuery ID of aws_glue_ml_transforms table (FK)|
|completed_on|timestamp without time zone|The last point in time that the requested task run was completed|
|error_string|text|The list of error strings associated with this task run|
|execution_time|bigint|The amount of time (in seconds) that the task run consumed resources|
|last_modified_on|timestamp without time zone|The last point in time that the requested task run was updated|
|log_group_name|text|The names of the log group for secure logging, associated with this task run|
|export_labels_task_run_properties_output_s3_path|text|The Amazon Simple Storage Service (Amazon S3) path where you will export the labels|
|find_matches_task_run_properties_job_id|text|The job ID for the Find Matches task run|
|find_matches_task_run_properties_job_name|text|The name assigned to the job for the Find Matches task run|
|find_matches_task_run_properties_job_run_id|text|The job run ID for the Find Matches task run|
|import_labels_task_run_properties_input_s3_path|text|The Amazon Simple Storage Service (Amazon S3) path from where you will import the labels|
|import_labels_task_run_properties_replace|boolean|Indicates whether to overwrite your existing labels|
|labeling_set_generation_task_run_properties_output_s3_path|text|The Amazon Simple Storage Service (Amazon S3) path where you will generate the labeling set|
|task_type|text|The type of task run|
|started_on|timestamp without time zone|The date and time that this task run started|
|status|text|The current status of the requested task run|
|id|text|The unique identifier for this task run|
|transform_id|text|The unique identifier for the transform|
