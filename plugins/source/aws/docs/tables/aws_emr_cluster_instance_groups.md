# Table: aws_emr_cluster_instance_groups

https://docs.aws.amazon.com/emr/latest/APIReference/API_InstanceGroup.html

The composite primary key for this table is (**cluster_arn**, **id**).

## Relations

This table depends on [aws_emr_clusters](aws_emr_clusters.md).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|cluster_arn (PK)|String|
|auto_scaling_policy|JSON|
|bid_price|String|
|configurations|JSON|
|configurations_version|Int|
|custom_ami_id|String|
|ebs_block_devices|JSON|
|ebs_optimized|Bool|
|id (PK)|String|
|instance_group_type|String|
|instance_type|String|
|last_successfully_applied_configurations|JSON|
|last_successfully_applied_configurations_version|Int|
|market|String|
|name|String|
|requested_instance_count|Int|
|running_instance_count|Int|
|shrink_policy|JSON|
|status|JSON|