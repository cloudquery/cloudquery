
# Table: aws_emr_clusters
The detailed description of the cluster.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text|The AWS Account ID of the resource.|
|region|text|The AWS Region of the resource.|
|applications|jsonb|The applications installed on this cluster.|
|auto_scaling_role|text|An IAM role for automatic scaling policies|
|auto_terminate|boolean|Specifies whether the cluster should terminate after completing all steps.|
|arn|text|The Amazon Resource Name of the cluster.|
|configurations|jsonb|The list of Configurations supplied to the EMR cluster.|
|custom_ami_id|text|The ID of a custom Amazon EBS-backed Linux AMI if the cluster uses a custom AMI.|
|ebs_root_volume_size|integer|The size, in GiB, of the Amazon EBS root device volume of the Linux AMI that is used for each EC2 instance|
|ec2_instance_attribute_additional_master_security_groups|text[]|A list of additional Amazon EC2 security group IDs for the master node.|
|ec2_instance_attribute_additional_slave_security_groups|text[]|A list of additional Amazon EC2 security group IDs for the core and task nodes.|
|ec2_instance_attribute_availability_zone|text|The Availability Zone in which the cluster will run.|
|ec2_instance_attribute_key_name|text|The name of the Amazon EC2 key pair to use when connecting with SSH into the master node as a user named "hadoop".|
|ec2_instance_attribute_subnet_id|text|Set this parameter to the identifier of the Amazon VPC subnet where you want the cluster to launch|
|ec2_instance_attribute_emr_managed_master_security_group|text|The identifier of the Amazon EC2 security group for the master node.|
|ec2_instance_attribute_emr_managed_slave_security_group|text|The identifier of the Amazon EC2 security group for the core and task nodes.|
|ec2_instance_attribute_iam_instance_profile|text|The IAM role that was specified when the cluster was launched|
|ec2_instance_attribute_requested_availability_zones|text[]|Specifies one or more Availability Zones in which to launch EC2 cluster instances when the EC2-Classic network configuration is supported.|
|ec2_instance_attribute_requested_subnet_ids|text[]|Specifies the unique identifier of one or more Amazon EC2 subnets in which to launch EC2 cluster instances.|
|ec2_instance_attribute_service_access_security_group|text|The identifier of the Amazon EC2 security group for the Amazon EMR service to access clusters in VPC private subnets.|
|id|text|The unique identifier for the cluster.|
|instance_collection_type|text|The instance group configuration of the cluster.|
|kerberos_kdc_admin_password|text|The password used within the cluster for the kadmin service on the cluster-dedicated KDC, which maintains Kerberos principals, password policies, and keytabs for the cluster. |
|kerberos_realm|text|The name of the Kerberos realm to which all nodes in a cluster belong|
|kerberos_ad_domain_join_password|text|The Active Directory password for ADDomainJoinUser.|
|kerberos_ad_domain_join_user|text|Required only when establishing a cross-realm trust with an Active Directory domain|
|kerberos_cross_realm_trust_principal_password|text|Required only when establishing a cross-realm trust with a KDC in a different realm|
|log_encryption_kms_key_id|text|The AWS KMS customer master key (CMK) used for encrypting log files|
|log_uri|text|The path to the Amazon S3 location where logs for this cluster are stored.|
|master_public_dns_name|text|The DNS name of the master node|
|name|text|The name of the cluster.|
|normalized_instance_hours|integer|An approximation of the cost of the cluster, represented in m1.small/hours|
|outpost_arn|text|The Amazon Resource Name (ARN) of the Outpost where the cluster is launched.|
|placement_groups|jsonb|Placement group configured for an Amazon EMR cluster.|
|release_label|text|The Amazon EMR release label, which determines the version of open-source application packages installed on the cluster|
|repo_upgrade_on_boot|text|Specifies the type of updates that are applied from the Amazon Linux AMI package repositories when an instance boots using the AMI.|
|requested_ami_version|text|The AMI version requested for this cluster.|
|running_ami_version|text|The AMI version running on this cluster.|
|scale_down_behavior|text|The way that individual Amazon EC2 instances terminate when an automatic scale-in activity occurs or an instance group is resized.|
|security_configuration|text|The name of the security configuration applied to the cluster.|
|service_role|text|The IAM role that will be assumed by the Amazon EMR service to access AWS resources on your behalf.|
|state|text|The current state of the cluster.|
|state_change_reason_code|text|The programmatic code for the state change reason.|
|state_change_reason_message|text|The descriptive message for the state change reason.|
|creation_date_time|timestamp without time zone|The creation date and time of the cluster.|
|end_date_time|timestamp without time zone|The date and time when the cluster was terminated.|
|ready_date_time|timestamp without time zone|The date and time when the cluster was ready to run steps.|
|step_concurrency_level|integer|Specifies the number of steps that can be executed concurrently.|
|tags|jsonb|A list of tags associated with a cluster.|
|termination_protected|boolean|Indicates whether Amazon EMR will lock the cluster to prevent the EC2 instances from being terminated by an API call or user intervention, or in the event of a cluster error.|
|visible_to_all_users|boolean|Indicates whether the cluster is visible to all IAM users of the AWS account associated with the cluster|
