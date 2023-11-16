# Table: aws_emr_steps

This table shows data for Amazon EMR Steps.

https://docs.aws.amazon.com/emr/latest/APIReference/API_Step.html

The composite primary key for this table is (**cluster_arn**, **id**).

## Relations

This table depends on [aws_emr_clusters](aws_emr_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn (PK)|`utf8`|
|action_on_failure|`utf8`|
|config|`json`|
|execution_role_arn|`utf8`|
|id (PK)|`utf8`|
|name|`utf8`|
|status|`json`|