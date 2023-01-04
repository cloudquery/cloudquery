# Table: aws_autoscaling_groups

https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_AutoScalingGroup.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_autoscaling_groups:
  - [aws_autoscaling_group_scaling_policies](aws_autoscaling_group_scaling_policies.md)
  - [aws_autoscaling_group_lifecycle_hooks](aws_autoscaling_group_lifecycle_hooks.md)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|String|
|_cq_sync_time|Timestamp|
|_cq_id|UUID|
|_cq_parent_id|UUID|
|account_id|String|
|region|String|
|load_balancers|JSON|
|load_balancer_target_groups|JSON|
|arn (PK)|String|
|auto_scaling_group|JSON|
|notification_configurations|JSON|