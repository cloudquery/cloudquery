
# Table: aws_glue_jobs
Specifies a job definition
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|arn|text|The Amazon Resource Name (ARN) of the workflow.|
|tags|jsonb|Resource tags.|
|allocated_capacity|bigint|This field is deprecated|
|code_gen_configuration_nodes|jsonb|The representation of a directed acyclic graph on which both the Glue Studio visual component and Glue Studio code generation is based|
|command_name|text|The name of the job command|
|command_python_version|text|The Python version being used to run a Python shell job|
|command_script_location|text|Specifies the Amazon Simple Storage Service (Amazon S3) path to a script that runs a job|
|connections|text[]|A list of connections used by the job|
|created_on|timestamp without time zone|The time and date that this job definition was created|
|default_arguments|jsonb|The default arguments for this job, specified as name-value pairs|
|description|text|A description of the job|
|execution_property_max_concurrent_runs|bigint|The maximum number of concurrent runs allowed for the job|
|glue_version|text|Glue version determines the versions of Apache Spark and Python that Glue supports|
|last_modified_on|timestamp without time zone|The last point in time when this job definition was modified|
|log_uri|text|This field is reserved for future use|
|max_capacity|float|For Glue version 10 or earlier jobs, using the standard worker type, the number of Glue data processing units (DPUs) that can be allocated when this job runs|
|max_retries|bigint|The maximum number of times to retry this job after a JobRun fails|
|name|text|The name you assign to this job definition|
|non_overridable_arguments|jsonb|Non-overridable arguments for this job, specified as name-value pairs|
|notification_property_notify_delay_after|bigint|After a job run starts, the number of minutes to wait before sending a job run delay notification|
|number_of_workers|bigint|The number of workers of a defined workerType that are allocated when a job runs|
|role|text|The name or Amazon Resource Name (ARN) of the IAM role associated with this job|
|security_configuration|text|The name of the SecurityConfiguration structure to be used with this job|
|timeout|bigint|The job timeout in minutes|
|worker_type|text|The type of predefined worker that is allocated when a job runs|
