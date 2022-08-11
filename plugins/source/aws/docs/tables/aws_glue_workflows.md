
# Table: aws_glue_workflows
A workflow is a collection of multiple dependent Glue jobs and crawlers that are run to complete a complex ETL task
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the workflow.|
|tags|jsonb|Resource tags.|
|blueprint_name|text|The name of the blueprint.|
|blueprint_run_id|text|The run ID for this blueprint.|
|created_on|timestamp without time zone|The date and time when the workflow was created.|
|default_run_properties|jsonb|A collection of properties to be used as part of each execution of the workflow. The run properties are made available to each job in the workflow|
|description|text|A description of the workflow.|
|last_modified_on|timestamp without time zone|The date and time when the workflow was last modified.|
|last_run_completed_on|timestamp without time zone|The date and time when the workflow run completed.|
|last_run_error_message|text|This error message describes any error that may have occurred in starting the workflow run|
|last_run_name|text|Name of the workflow that was run.|
|last_run_previous_run_id|text|The ID of the previous workflow run.|
|last_run_started_on|timestamp without time zone|The date and time when the workflow run was started.|
|last_run_starting_event_batch_condition_size|bigint|Number of events in the batch.|
|last_run_starting_event_batch_condition_window|bigint|Duration of the batch window in seconds.|
|last_run_statistics_failed_actions|bigint|Total number of Actions that have failed.|
|last_run_statistics_running_actions|bigint|Total number Actions in running state.|
|last_run_statistics_stopped_actions|bigint|Total number of Actions that have stopped.|
|last_run_statistics_succeeded_actions|bigint|Total number of Actions that have succeeded.|
|last_run_statistics_timeout_actions|bigint|Total number of Actions that timed out.|
|last_run_statistics_total_actions|bigint|Total number of Actions in the workflow run.|
|last_run_status|text|The status of the workflow run.|
|last_run_workflow_run_id|text|The ID of this workflow run.|
|last_run_workflow_run_properties|jsonb|The workflow run properties which were set during the run.|
|max_concurrent_runs|bigint|You can use this parameter to prevent unwanted multiple updates to data, to control costs, or in some cases, to prevent exceeding the maximum number of concurrent runs of any of the component jobs|
|name|text|The name of the workflow.|
