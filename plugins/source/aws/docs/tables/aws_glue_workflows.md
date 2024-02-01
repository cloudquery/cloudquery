# Table: aws_glue_workflows

This table shows data for Glue Workflows.

https://docs.aws.amazon.com/glue/latest/webapi/API_Workflow.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn|`utf8`|
|tags|`json`|
|blueprint_details|`json`|
|created_on|`timestamp[us, tz=UTC]`|
|default_run_properties|`json`|
|description|`utf8`|
|graph|`json`|
|last_modified_on|`timestamp[us, tz=UTC]`|
|last_run|`json`|
|max_concurrent_runs|`int64`|
|name|`utf8`|