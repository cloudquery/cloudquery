# Table: aws_detective_graphs

This table shows data for Amazon Detective Graphs.

https://docs.aws.amazon.com/detective/latest/APIReference/API_ListGraphs.html

The primary key for this table is **_cq_id**.
The following field is used to calculate the value of `_cq_id`: **arn**.
## Relations

The following tables depend on aws_detective_graphs:
  - [aws_detective_graph_members](aws_detective_graph_members.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|