# Table: aws_emr_notebook_executions

This table shows data for Amazon EMR Notebook Executions.

https://docs.aws.amazon.com/emr/latest/APIReference/API_NotebookExecution.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_emr_clusters](aws_emr_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn|`utf8`|
|arn (PK)|`utf8`|
|editor_id|`utf8`|
|end_time|`timestamp[us, tz=UTC]`|
|environment_variables|`json`|
|execution_engine|`json`|
|last_state_change_reason|`utf8`|
|notebook_execution_id|`utf8`|
|notebook_execution_name|`utf8`|
|notebook_instance_security_group_id|`utf8`|
|notebook_params|`utf8`|
|notebook_s3_location|`json`|
|output_notebook_format|`utf8`|
|output_notebook_s3_location|`json`|
|output_notebook_uri|`utf8`|
|start_time|`timestamp[us, tz=UTC]`|
|status|`utf8`|
|tags|`json`|