# Table: aws_emr_clusters

https://docs.aws.amazon.com/emr/latest/APIReference/API_Cluster.html

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
|applications|JSON|
|auto_scaling_role|String|
|auto_terminate|Bool|
|configurations|JSON|
|custom_ami_id|String|
|ebs_root_volume_size|Int|
|ec2_instance_attributes|JSON|
|id|String|
|instance_collection_type|String|
|kerberos_attributes|JSON|
|log_encryption_kms_key_id|String|
|log_uri|String|
|master_public_dns_name|String|
|name|String|
|normalized_instance_hours|Int|
|os_release_label|String|
|outpost_arn|String|
|placement_groups|JSON|
|release_label|String|
|repo_upgrade_on_boot|String|
|requested_ami_version|String|
|running_ami_version|String|
|scale_down_behavior|String|
|security_configuration|String|
|service_role|String|
|status|JSON|
|step_concurrency_level|Int|
|tags|JSON|
|termination_protected|Bool|
|visible_to_all_users|Bool|