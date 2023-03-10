# Table: aws_eks_fargate_profiles

This table shows data for AWS EKS Fargate Profiles.

https://docs.aws.amazon.com/eks/latest/APIReference/API_FargateProfile.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_eks_clusters](aws_eks_clusters).

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
|cluster_name|String|
|created_at|Timestamp|
|fargate_profile_arn|String|
|fargate_profile_name|String|
|pod_execution_role_arn|String|
|selectors|JSON|
|status|String|
|subnets|StringArray|
|tags|JSON|