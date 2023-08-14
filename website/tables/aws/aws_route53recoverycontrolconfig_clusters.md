# Table: aws_route53recoverycontrolconfig_clusters

This table shows data for Amazon Route 53 Application Recovery Controller Recovery Control Configuration Clusters.

https://docs.aws.amazon.com/recovery-cluster/latest/api/cluster.html

The primary key for this table is **arn**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|arn (PK)|`utf8`|
|cluster_arn|`utf8`|
|cluster_endpoints|`json`|
|name|`utf8`|
|status|`utf8`|