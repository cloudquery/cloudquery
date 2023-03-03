# Table: aws_rds_cluster_backtracks

https://docs.aws.amazon.com/AmazonRDS/latest/APIReference/API_DescribeDBClusterBacktracks.html

The composite primary key for this table is (**db_cluster_arn**, **backtrack_identifier**).

## Relations

This table depends on [aws_rds_clusters](aws_rds_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|db_cluster_arn (PK)|String|
|backtrack_identifier (PK)|String|
|backtrack_request_creation_time|Timestamp|
|backtrack_to|Timestamp|
|backtracked_from|Timestamp|
|db_cluster_identifier|String|
|status|String|