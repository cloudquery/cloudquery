# Table: aws_emr_clusters

This table shows data for Amazon EMR Clusters.

https://docs.aws.amazon.com/emr/latest/APIReference/API_Cluster.html

The primary key for this table is **arn**.

## Relations

The following tables depend on aws_emr_clusters:
  - [aws_emr_cluster_instance_fleets](aws_emr_cluster_instance_fleets)
  - [aws_emr_cluster_instance_groups](aws_emr_cluster_instance_groups)
  - [aws_emr_cluster_instances](aws_emr_cluster_instances)

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id|`utf8`|
|region|`utf8`|
|arn (PK)|`utf8`|
|tags|`json`|
|applications|`json`|
|auto_scaling_role|`utf8`|
|auto_terminate|`bool`|
|cluster_arn|`utf8`|
|configurations|`json`|
|custom_ami_id|`utf8`|
|ebs_root_volume_size|`int64`|
|ec2_instance_attributes|`json`|
|id|`utf8`|
|instance_collection_type|`utf8`|
|kerberos_attributes|`json`|
|log_encryption_kms_key_id|`utf8`|
|log_uri|`utf8`|
|master_public_dns_name|`utf8`|
|name|`utf8`|
|normalized_instance_hours|`int64`|
|os_release_label|`utf8`|
|outpost_arn|`utf8`|
|placement_groups|`json`|
|release_label|`utf8`|
|repo_upgrade_on_boot|`utf8`|
|requested_ami_version|`utf8`|
|running_ami_version|`utf8`|
|scale_down_behavior|`utf8`|
|security_configuration|`utf8`|
|service_role|`utf8`|
|status|`json`|
|step_concurrency_level|`int64`|
|termination_protected|`bool`|
|visible_to_all_users|`bool`|

## Example Queries

These SQL queries are sampled from CloudQuery policies and are compatible with PostgreSQL.

### EMR clusters should not have public IP addresses

```sql
SELECT
  'EMR clusters should not have public IP addresses' AS title,
  aws_emr_clusters.account_id,
  aws_emr_clusters.arn AS resource_id,
  CASE
  WHEN aws_ec2_subnets.map_public_ip_on_launch
  AND aws_emr_clusters.status->>'State' IN ('RUNNING', 'WAITING')
  THEN 'fail'
  ELSE 'pass'
  END
    AS status
FROM
  aws_emr_clusters
  LEFT JOIN aws_ec2_subnets ON
      aws_emr_clusters.ec2_instance_attributes->>'Ec2SubnetId'
      = aws_ec2_subnets.subnet_id;
```


