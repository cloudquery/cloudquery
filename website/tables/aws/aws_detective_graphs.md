# Table: aws_detective_graphs

This table shows data for Amazon Detective Graphs.

https://docs.aws.amazon.com/detective/latest/APIReference/API_ListGraphs.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_detective_graphs:
  - [aws_detective_graph_members](aws_detective_graph_members)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|tags|`json`|
|arn (PK)|`utf8`|
|created_time|`timestamp[us, tz=UTC]`|