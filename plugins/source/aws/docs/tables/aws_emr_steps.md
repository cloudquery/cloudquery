# Table: aws_emr_steps

This table shows data for Amazon EMR Steps.

https://docs.aws.amazon.com/emr/latest/APIReference/API_Step.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**cluster_arn**, **id**).
## Relations

This table depends on [aws_emr_clusters](aws_emr_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn|`utf8`|
|action_on_failure|`utf8`|
|config|`json`|
|execution_role_arn|`utf8`|
|id|`utf8`|
|name|`utf8`|
|status|`json`|