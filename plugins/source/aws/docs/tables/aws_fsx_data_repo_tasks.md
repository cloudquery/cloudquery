
# Table: aws_fsx_data_repo_tasks
A description of the data repository task
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|creation_time|timestamp without time zone|The time that the resource was created, in seconds (since 1970-01-01T00:00:00Z), also known as Unix time|
|file_system_id|text|The globally unique ID of the file system, assigned by Amazon FSx|
|lifecycle|text|The lifecycle status of the data repository task, as follows:  * PENDING - Amazon FSx has not started the task  * EXECUTING - Amazon FSx is processing the task  * FAILED - Amazon FSx was not able to complete the task|
|task_id|text|The system-generated, unique 17-digit ID of the data repository task|
|type|text|The type of data repository task  * The EXPORT_TO_REPOSITORY data repository task exports from your Lustre file system from to a linked S3 bucket  * The IMPORT_METADATA_FROM_REPOSITORY data repository task imports metadata changes from a linked S3 bucket to your Lustre file system|
|end_time|timestamp without time zone|The time that Amazon FSx completed processing the task, populated after the task is complete|
|failure_details_message|text|A detailed error message|
|paths|text[]|An array of paths on the Amazon FSx for Lustre file system that specify the data for the data repository task to process|
|report_enabled|boolean|Set Enabled to True to generate a CompletionReport when the task completes|
|report_format|text|Required if Enabled is set to true|
|report_path|text|Required if Enabled is set to true|
|report_scope|text|Required if Enabled is set to true|
|arn|text|The Amazon Resource Name (ARN) for a given resource|
|start_time|timestamp without time zone|The time that Amazon FSx began processing the task|
|status_failed_count|bigint|A running total of the number of files that the task failed to process|
|status_last_updated_time|timestamp without time zone|The time at which the task status was last updated|
|status_succeeded_count|bigint|A running total of the number of files that the task has successfully processed|
|status_total_count|bigint|The total number of files that the task will process|
|tags|jsonb|A list of Tag values, with a maximum of 50 elements|
