# Table: aws_eks_cluster_node_groups

https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_eks_clusters](aws_eks_clusters.md).

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
|ami_type|String|
|capacity_type|String|
|cluster_name|String|
|created_at|Timestamp|
|disk_size|Int|
|health|JSON|
|instance_types|StringArray|
|labels|JSON|
|launch_template|JSON|
|modified_at|Timestamp|
|node_role|String|
|nodegroup_arn|String|
|nodegroup_name|String|
|release_version|String|
|remote_access|JSON|
|resources|JSON|
|scaling_config|JSON|
|status|String|
|subnets|StringArray|
|tags|JSON|
|taints|JSON|
|update_config|JSON|
|version|String|