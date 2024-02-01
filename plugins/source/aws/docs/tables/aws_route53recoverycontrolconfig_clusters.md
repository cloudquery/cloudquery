# Table: aws_route53recoverycontrolconfig_clusters

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Control Configuration Clusters.

https://docs.aws.amazon.com/recovery-cluster/latest/api/cluster.html

The primary key for this table is **_cq_id**.
The following fields are used to calculate the value of `_cq_id`: (**request_account_id**, **arn**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|request_account_id|`utf8`|
|arn|`utf8`|
|cluster_arn|`utf8`|
|cluster_endpoints|`json`|
|name|`utf8`|
|owner|`utf8`|
|status|`utf8`|