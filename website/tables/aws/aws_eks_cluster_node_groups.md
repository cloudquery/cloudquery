# Table: aws_eks_cluster_node_groups

This table shows data for Amazon Elastic Kubernetes Service (EKS) Cluster Node Groups.

https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_eks_clusters](aws_eks_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|ami_type|`utf8`|
|capacity_type|`utf8`|
|cluster_name|`utf8`|
|created_at|`timestamp[us, tz=UTC]`|
|disk_size|`int64`|
|health|`json`|
|instance_types|`list<item: utf8, nullable>`|
|labels|`json`|
|launch_template|`json`|
|modified_at|`timestamp[us, tz=UTC]`|
|node_role|`utf8`|
|nodegroup_arn|`utf8`|
|nodegroup_name|`utf8`|
|release_version|`utf8`|
|remote_access|`json`|
|resources|`json`|
|scaling_config|`json`|
|status|`utf8`|
|subnets|`list<item: utf8, nullable>`|
|tags|`json`|
|taints|`json`|
|update_config|`json`|
|version|`utf8`|