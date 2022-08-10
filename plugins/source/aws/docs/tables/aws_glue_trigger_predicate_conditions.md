
# Table: aws_glue_trigger_predicate_conditions
Defines a condition under which a trigger fires
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|trigger_cq_id|uuid|Unique CloudQuery ID of aws_glue_triggers table (FK)|
|crawl_state|text|The state of the crawler to which this condition applies|
|crawler_name|text|The name of the crawler to which this condition applies|
|job_name|text|The name of the job whose JobRuns this condition applies to, and on which this trigger waits|
|logical_operator|text|A logical operator|
|state|text|The condition state|
