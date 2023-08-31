# Table: aws_emr_cluster_instances

This table shows data for Amazon EMR Cluster Instances.

https://docs.aws.amazon.com/emr/latest/APIReference/API_Instance.html

The primary key for this table is **arn**.

## Relations

This table depends on [aws_emr_clusters](aws_emr_clusters).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|ebs_volumes|`json`|
|ec2_instance_id|`utf8`|
|id|`utf8`|
|instance_fleet_id|`utf8`|
|instance_group_id|`utf8`|
|instance_type|`utf8`|
|market|`utf8`|
|private_dns_name|`utf8`|
|private_ip_address|`utf8`|
|public_dns_name|`utf8`|
|public_ip_address|`utf8`|
|status|`json`|