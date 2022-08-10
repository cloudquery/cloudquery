
# Table: aws_glue_job_runs
Contains information about a job run
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|job_cq_id|uuid|Unique CloudQuery ID of aws_glue_jobs table (FK)|
|allocated_capacity|bigint|This field is deprecated|
|arguments|jsonb|The job arguments associated with this run|
|attempt|bigint|The number of the attempt to run this job|
|completed_on|timestamp without time zone|The date and time that this job run completed|
|dpu_seconds|float|This field populates only for Auto Scaling job runs, and represents the total time each executor ran during the lifecycle of a job run in seconds, multiplied by a DPU factor (1 for G1X, 2 for G2X, or 025 for G025X workers)|
|error_message|text|An error message associated with this job run|
|execution_time|bigint|The amount of time (in seconds) that the job run consumed resources|
|glue_version|text|Glue version determines the versions of Apache Spark and Python that Glue supports|
|id|text|The ID of this job run|
|job_name|text|The name of the job definition being used in this run|
|job_run_state|text|The current state of the job run|
|last_modified_on|timestamp without time zone|The last time that this job run was modified|
|log_group_name|text|The name of the log group for secure logging that can be server-side encrypted in Amazon CloudWatch using KMS|
|max_capacity|float|The number of Glue data processing units (DPUs) that can be allocated when this job runs|
|notification_property_notify_delay_after|bigint|After a job run starts, the number of minutes to wait before sending a job run delay notification|
|number_of_workers|bigint|The number of workers of a defined workerType that are allocated when a job runs|
|predecessor_runs|jsonb|A list of predecessors to this job run|
|previous_run_id|text|The ID of the previous run of this job|
|security_configuration|text|The name of the SecurityConfiguration structure to be used with this job run|
|started_on|timestamp without time zone|The date and time at which this job run was started|
|timeout|bigint|The JobRun timeout in minutes|
|trigger_name|text|The name of the trigger that started this job run|
|worker_type|text|The type of predefined worker that is allocated when a job runs|
