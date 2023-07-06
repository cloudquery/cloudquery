# Table: aws_rds_cluster_backtracks

This table shows data for Amazon Relational Database Service (RDS) Cluster Backtracks.

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeDBClusterBacktracks.html

The composite primary key for this table is (**db_cluster_arn**, **backtrack_identifier**).

## Relations

This table depends on [aws_rds_clusters](aws_rds_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|db_cluster_arn (PK)|`utf8`|
|backtrack_identifier (PK)|`utf8`|
|backtrack_request_creation_time|`timestamp[us, tz=UTC]`|
|backtrack_to|`timestamp[us, tz=UTC]`|
|backtracked_from|`timestamp[us, tz=UTC]`|
|db_cluster_identifier|`utf8`|
|status|`utf8`|