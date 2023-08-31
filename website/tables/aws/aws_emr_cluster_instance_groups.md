# Table: aws_emr_cluster_instance_groups

This table shows data for Amazon EMR Cluster Instance Groups.

https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceGroup.html

The composite primary key for this table is (**cluster_arn**, **id**).

## Relations

This table depends on [aws_emr_clusters](aws_emr_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|cluster_arn (PK)|`utf8`|
|auto_scaling_policy|`json`|
|bid_price|`utf8`|
|configurations|`json`|
|configurations_version|`int64`|
|custom_ami_id|`utf8`|
|ebs_block_devices|`json`|
|ebs_optimized|`bool`|
|id (PK)|`utf8`|
|instance_group_type|`utf8`|
|instance_type|`utf8`|
|last_successfully_applied_configurations|`json`|
|last_successfully_applied_configurations_version|`int64`|
|market|`utf8`|
|name|`utf8`|
|requested_instance_count|`int64`|
|running_instance_count|`int64`|
|shrink_policy|`json`|
|status|`json`|