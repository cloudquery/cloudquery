# Table: aws_redshift_subnet_groups

https://docs.aws.amazon.com/redshift/latest/APIReference/API_ClusterSubnetGroup.html

The primary key for this table is **arn**.



## Columns
| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|arn (PK)|String|
|cluster_subnet_group_name|String|
|description|String|
|subnet_group_status|String|
|subnets|JSON|
|tags|JSON|
|vpc_id|String|