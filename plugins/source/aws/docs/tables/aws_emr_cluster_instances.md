# Table: aws_emr_cluster_instances

https://docs.aws.amazon.com/emr/latest/APIReference/API_Instance.html

The primary key for this table is **arn**.

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
|arn (PK)|String|
|ebs_volumes|JSON|
|ec2_instance_id|String|
|id|String|
|instance_fleet_id|String|
|instance_group_id|String|
|instance_type|String|
|market|String|
|private_dns_name|String|
|private_ip_address|String|
|public_dns_name|String|
|public_ip_address|String|
|status|JSON|