
# Table: aws_glue_trigger_actions
Defines an action to be initiated by a trigger
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|trigger_cq_id|uuid|Unique CloudQuery ID of aws_glue_triggers table (FK)|
|arguments|jsonb|The job arguments used when this trigger fires|
|crawler_name|text|The name of the crawler to be used with this action|
|job_name|text|The name of a job to be run|
|notify_delay_after|bigint|After a job run starts, the number of minutes to wait before sending a job run delay notification|
|security_configuration|text|The name of the SecurityConfiguration structure to be used with this action|
|timeout|bigint|The JobRun timeout in minutes|
