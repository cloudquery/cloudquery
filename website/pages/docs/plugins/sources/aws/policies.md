# Policies and Compliance Frameworks

Policies are a set of standard SQL queries that can be run to check the security and compliance of your cloud resources against best practice frameworks.

This page documents the available CloudQuery SQL Policies for AWS. See the [readme on GitHub](https://github.com/cloudquery/cloudquery/tree/main/plugins/source/aws/policies) for installation instructions.
## AWS CIS V1.2.0

### Requirements
AWS CIS V1.2.0 requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - aws_cloudtrail_trail_event_selectors
  - aws_cloudtrail_trails
  - aws_cloudwatch_alarms
  - aws_cloudwatchlogs_metric_filters
  - aws_ec2_flow_logs
  - aws_ec2_security_groups
  - aws_ec2_vpcs
  - aws_iam_credential_reports
  - aws_iam_password_policies
  - aws_iam_user_access_keys
  - aws_iam_users
  - aws_iam_virtual_mfa_devices
  - aws_kms_keys
  - aws_s3_buckets
  - aws_sns_subscriptions
```

### Queries

AWS CIS V1.2.0 performs the following checks:
  - Avoid the use of "root" account. Show used in last 30 days (Scored)
  - Ensure MFA is enabled for all IAM users that have a console password (Scored)
  - Ensure credentials unused for 90 days or greater are disabled (Scored)
  - Ensure access keys are rotated every 90 days or less
  - Ensure IAM password policy requires at least one uppercase letter
  - Ensure IAM password policy requires at least one lowercase letter
  - Ensure IAM password policy requires at least one symbol
  - Ensure IAM password policy requires at least one number
  - Ensure IAM password policy requires minimum length of 14 or greater
  - Ensure IAM password policy prevents password reuse
  - Ensure IAM password policy expires passwords within 90 days or less
  - Ensure no root account access key exists (Scored)
  - Ensure MFA is enabled for the "root" account
  - Ensure hardware MFA is enabled for the "root" account (Scored)
  - Ensure CloudTrail is enabled in all regions
  - Ensure CloudTrail log file validation is enabled
  - CloudTrail trails should be integrated with CloudWatch Logs
  - Ensure S3 bucket access logging is enabled on the CloudTrail S3 bucket
  - CloudTrail should have encryption at rest enabled
  - Ensure rotation for customer created custom master keys is enabled (Scored)
  - VPC flow logging should be enabled in all VPCs
  - Ensure a log metric filter and alarm exist for Management Console sign-in without MFA (Scored)
  - Ensure a log metric filter and alarm exist for usage of "root" account (Score)
  - Ensure a log metric filter and alarm exist for IAM policy changes (Score)
  - Ensure a log metric filter and alarm exist for CloudTrail configuration changes (Scored)
  - Ensure a log metric filter and alarm exist for AWS Management Console authentication failures (Scored)
  - Ensure a log metric filter and alarm exist for disabling or scheduled deletion of customer created CMKs (Scored)
  - Ensure a log metric filter and alarm exist for S3 bucket policy changes (Scored)
  - Ensure a log metric filter and alarm exist for AWS Config configuration changes (Scored)
  - Ensure a log metric filter and alarm exist for security group changes (Scored)
  - Ensure a log metric filter and alarm exist for changes to Network Access Control Lists (NACL) (Scored)
  - Ensure a log metric filter and alarm exist for changes to network gateways (Scored)
  - Ensure a log metric filter and alarm exist for route table changes (Scored)
  - Ensure a log metric filter and alarm exist for VPC changes (Scored)
  - Ensure no security groups allow ingress from 0.0.0.0/0 to port 22 (Scored)
  - Ensure no security groups allow ingress from 0.0.0.0/0 to port 3389 (Scored)
  - The VPC default security group should not allow inbound and outbound traffic

### Dependent Views

AWS CIS V1.2.0 depends on the following views:

  - view_aws_log_metric_filter_and_alarm<sup>*</sup>
  - view_aws_security_group_ingress_rules<sup>*</sup>

  <sup>*</sup> These views are automatically created or updated by this policy.
## AWS PCI DSS v3.2.1

### Requirements
AWS PCI DSS v3.2.1 requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - aws_autoscaling_groups
  - aws_cloudtrail_trail_event_selectors
  - aws_cloudtrail_trails
  - aws_cloudwatch_alarms
  - aws_cloudwatchlogs_metric_filters
  - aws_codebuild_projects
  - aws_config_configuration_recorders
  - aws_dms_replication_instances
  - aws_ec2_ebs_snapshot_attributes
  - aws_ec2_ebs_snapshots
  - aws_ec2_eips
  - aws_ec2_flow_logs
  - aws_ec2_instances
  - aws_ec2_security_groups
  - aws_ec2_vpcs
  - aws_elasticsearch_domains
  - aws_elbv2_listeners
  - aws_elbv2_load_balancers
  - aws_guardduty_detectors
  - aws_iam_accounts
  - aws_iam_credential_reports
  - aws_iam_password_policies
  - aws_iam_policies
  - aws_iam_user_access_keys
  - aws_iam_user_attached_policies
  - aws_iam_users
  - aws_iam_virtual_mfa_devices
  - aws_kms_keys
  - aws_lambda_functions
  - aws_rds_cluster_snapshots
  - aws_rds_instances
  - aws_redshift_clusters
  - aws_regions
  - aws_s3_accounts
  - aws_s3_bucket_encryption_rules
  - aws_s3_bucket_grants
  - aws_s3_buckets
  - aws_sagemaker_notebook_instances
  - aws_secretsmanager_secrets
  - aws_sns_subscriptions
  - aws_ssm_instance_compliance_items
  - aws_ssm_instances
  - aws_waf_web_acls
  - aws_wafv2_web_acls
```

### Queries

AWS PCI DSS v3.2.1 performs the following checks:
  - Auto Scaling groups associated with a load balancer should use health checks
  - CloudTrail should have encryption at rest enabled
  - Ensure CloudTrail is enabled in all regions
  - Ensure CloudTrail log file validation is enabled
  - CloudTrail trails should be integrated with CloudWatch Logs
  - CodeBuild GitHub or Bitbucket source repository URLs should use OAuth
  - CodeBuild project environment variables should not contain clear text credentials
  - AWS Config should be enabled
  - Ensure a log metric filter and alarm exist for usage of "root" account (Score)
  - AWS Database Migration Service replication instances should not be public
  - Amazon EBS snapshots should not be public, determined by the ability to be restorable by anyone
  - The VPC default security group should not allow inbound and outbound traffic
  - Unused EC2 EIPs should be removed
  - Ensure no security groups allow ingress from 0.0.0.0/0 to port 22 (Scored)
  - VPC flow logging should be enabled in all VPCs
  - Application Load Balancer should be configured to redirect all HTTP requests to HTTPS
  - Elasticsearch domains should be in a VPC
  - Elasticsearch domains should have encryption at rest enabled
  - GuardDuty should be enabled
  - Ensure no root account access key exists (Scored)
  - IAM users should not have IAM policies attached
  - IAM policies should not allow full ''*'' administrative privileges
  - Ensure hardware MFA is enabled for the "root" account (Scored)
  - Ensure MFA is enabled for the "root" account
  - Ensure MFA is enabled for all IAM users that have a console password (Scored)
  - Ensure credentials unused for 90 days or greater are disabled (Scored)
  - Password policies for IAM users should have strong configurations
  - Ensure rotation for customer created custom master keys is enabled (Scored)
  - Lambda functions should prohibit public access
  - Lambda functions should be in a VPC
  - RDS snapshots should be private
  - RDS DB instances should prohibit public access, determined by the PubliclyAccessible configuration
  - Amazon Redshift clusters should prohibit public access
  - S3 buckets should prohibit public write access
  - S3 buckets should prohibit public read access
  - S3 buckets with replication rules should be enabled
  - S3 buckets should have server-side encryption enabled
  - S3 buckets should deny non-HTTPS requests
  - S3 Block Public Access setting should be enabled
  - Amazon SageMaker notebook instances should not have direct internet access
  - Secrets Manager secrets should have automatic rotation enabled
  - Secrets Manager secrets configured with automatic rotation should rotate successfully
  - Remove unused Secrets Manager secrets
  - Secrets Manager secrets should be rotated within a specified number of days
  - Amazon EC2 instances managed by Systems Manager should have a patch compliance status of COMPLIANT after a patch installation
  - Amazon EC2 instances managed by Systems Manager should have an association compliance status of COMPLIANT
  - Amazon EC2 instances should be managed by AWS Systems Manager
  - AWS WAF Classic global web ACL logging should be enabled

### Dependent Views

AWS PCI DSS v3.2.1 depends on the following views:

  - view_aws_log_metric_filter_and_alarm<sup>*</sup>
  - view_aws_security_group_ingress_rules<sup>*</sup>

  <sup>*</sup> These views are automatically created or updated by this policy.
## AWS Foundational Security Best Practices

### Requirements
AWS Foundational Security Best Practices requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - aws_account_alternate_contacts
  - aws_acm_certificates
  - aws_apigateway_rest_api_stages
  - aws_apigateway_rest_apis
  - aws_apigatewayv2_api_stages
  - aws_apigatewayv2_apis
  - aws_applicationautoscaling_policies
  - aws_autoscaling_groups
  - aws_cloudfront_distributions
  - aws_cloudtrail_trail_event_selectors
  - aws_cloudtrail_trails
  - aws_codebuild_projects
  - aws_config_configuration_recorders
  - aws_dax_clusters
  - aws_dms_replication_instances
  - aws_dynamodb_table_continuous_backups
  - aws_dynamodb_table_replica_auto_scalings
  - aws_dynamodb_tables
  - aws_ec2_ebs_snapshot_attributes
  - aws_ec2_ebs_snapshots
  - aws_ec2_ebs_volumes
  - aws_ec2_flow_logs
  - aws_ec2_instances
  - aws_ec2_network_acls
  - aws_ec2_regional_configs
  - aws_ec2_security_groups
  - aws_ec2_subnets
  - aws_ec2_vpc_endpoints
  - aws_ec2_vpcs
  - aws_ecs_cluster_services
  - aws_ecs_clusters
  - aws_ecs_task_definitions
  - aws_efs_filesystems
  - aws_elasticbeanstalk_configuration_settings
  - aws_elasticbeanstalk_environments
  - aws_elasticsearch_domains
  - aws_elbv1_load_balancers
  - aws_elbv2_listeners
  - aws_elbv2_load_balancer_attributes
  - aws_elbv2_load_balancers
  - aws_emr_clusters
  - aws_guardduty_detectors
  - aws_iam_accounts
  - aws_iam_credential_reports
  - aws_iam_group_policies
  - aws_iam_groups
  - aws_iam_password_policies
  - aws_iam_policies
  - aws_iam_role_policies
  - aws_iam_roles
  - aws_iam_user_access_keys
  - aws_iam_user_attached_policies
  - aws_iam_user_policies
  - aws_iam_users
  - aws_iam_virtual_mfa_devices
  - aws_kms_keys
  - aws_lambda_functions
  - aws_lambda_runtimes
  - aws_rds_cluster_snapshots
  - aws_rds_clusters
  - aws_rds_db_parameter_groups
  - aws_rds_db_security_groups
  - aws_rds_db_snapshots
  - aws_rds_event_subscriptions
  - aws_rds_instances
  - aws_redshift_cluster_parameter_groups
  - aws_redshift_cluster_parameters
  - aws_redshift_clusters
  - aws_regions
  - aws_s3_accounts
  - aws_s3_bucket_encryption_rules
  - aws_s3_bucket_grants
  - aws_s3_buckets
  - aws_sagemaker_notebook_instances
  - aws_secretsmanager_secrets
  - aws_sns_topics
  - aws_sqs_queues
  - aws_ssm_documents
  - aws_ssm_instance_compliance_items
  - aws_ssm_instances
  - aws_waf_web_acls
```

### Queries

AWS Foundational Security Best Practices performs the following checks:
  - Security contact information should be provided for an AWS account
  - certificate has less than 30 days to be renewed
  - API Gateway REST and WebSocket API logging should be enabled
  - API Gateway REST API stages should be configured to use SSL certificates for backend authentication
  - API Gateway REST API stages should have AWS X-Ray tracing enabled
  - API Gateway should be associated with an AWS WAF web ACL
  - API Gateway REST API cache data should be encrypted at rest
  - Auto Scaling groups associated with a load balancer should use health checks
  - AWS Config should be enabled
  - CloudFront distributions should have a default root object configured
  - CloudFront distributions should have origin access identity enabled
  - CloudFront distributions should require encryption in transit
  - CloudFront distributions should have origin failover configured
  - CloudFront distributions should have logging enabled
  - API Gateway should be associated with an AWS WAF web ACL
  - Ensure CloudTrail is enabled in all regions
  - CloudTrail should have encryption at rest enabled
  - Ensure CloudTrail log file validation is enabled
  - CloudTrail trails should be integrated with CloudWatch Logs
  - CodeBuild GitHub or Bitbucket source repository URLs should use OAuth
  - CodeBuild project environment variables should not contain clear text credentials
  - AWS Database Migration Service replication instances should not be public
  - DynamoDB tables should automatically scale capacity with demand
  - DynamoDB tables should have point-in-time recovery enabled
  - DynamoDB Accelerator (DAX) clusters should be encrypted at rest
  - Amazon EBS snapshots should not be public, determined by the ability to be restorable by anyone
  - The VPC default security group should not allow inbound and outbound traffic
  - Attached EBS volumes should be encrypted at rest
  - Stopped EC2 instances should be removed after a specified time period
  - VPC flow logging should be enabled in all VPCs
  - EBS default encryption should be enabled
  - EC2 instances should use IMDSv2
  - EC2 instances should not have a public IP address
  - Amazon EC2 should be configured to use VPC endpoints that are created for the Amazon EC2 service
  - EC2 subnets should not automatically assign public IP addresses
  - Unused network access control lists should be removed
  - EC2 instances should not use multiple ENIs
  - Aggregates rules of security groups with ports and IPs including ipv6
  - Security groups should not allow unrestricted access to ports with high risk
  - Amazon ECS task definitions should have secure networking modes and user definitions
  - Amazon ECS services should not have public IP addresses assigned to them automatically
  - Amazon EFS should be configured to encrypt file data at rest using AWS KMS
  - Amazon EFS volumes should be in backup plans
  - Elastic Beanstalk environments should have enhanced health reporting enabled
  - Elastic Beanstalk managed platform updates should be enabled
  - Elasticsearch domains should have encryption at rest enabled
  - Elasticsearch domains should be in a VPC
  - Elasticsearch domains should encrypt data sent between nodes
  - Elasticsearch domain error logging to CloudWatch Logs should be enabled
  - Elasticsearch domains should have audit logging enabled
  - Elasticsearch domains should have at least three data nodes
  - Elasticsearch domains should be configured with at least three dedicated master nodes
  - Connections to Elasticsearch domains should be encrypted using TLS 1.2
  - Classic Load Balancers with SSL/HTTPS listeners should use a certificate provided by AWS Certificate Manager
  - Classic Load Balancer listeners should be configured with HTTPS or TLS termination
  - Application load balancers should be configured to drop HTTP headers
  - Application and Classic Load Balancers logging should be enabled
  - Application Load Balancer deletion protection should be enabled
  - Classic Load Balancers should have connection draining enabled
  - Classic Load Balancers with HTTPS/SSL listeners should use a predefined security policy that has strong configuration
  - Application Load Balancer should be configured to redirect all HTTP requests to HTTPS
  - EMR clusters should not have public IP addresses
  - GuardDuty should be enabled
  - IAM policies should not allow full ''*'' administrative privileges
  - IAM users should not have IAM policies attached
  - IAM users'' access keys should be rotated every 90 days or less
  - Ensure no root account access key exists (Scored)
  - Ensure MFA is enabled for all IAM users that have a console password (Scored)
  - Ensure hardware MFA is enabled for the "root" account (Scored)
  - Password policies for IAM users should have strong configurations
  - Unused IAM user credentials should be removed
  - IAM customer managed policies that you create should not allow wildcard actions for services
  - IAM customer managed policies should not allow decryption and re-encryption actions on all KMS keys
  - IAM principals should not have IAM inline policies that allow decryption and re-encryption actions on all KMS keys
  - AWS KMS keys should not be unintentionally deleted
  - Lambda functions should prohibit public access
  - Lambda functions should use supported runtimes
  - RDS snapshots should be private
  - RDS DB instances should prohibit public access, determined by the PubliclyAccessible configuration
  - RDS DB instances should have encryption at rest enabled
  - RDS cluster snapshots and database snapshots should be encrypted at rest
  - RDS DB instances should be configured with multiple Availability Zones
  - Enhanced monitoring should be configured for RDS DB instances and clusters
  - RDS clusters should have deletion protection enabled
  - RDS DB instances should have deletion protection enabled
  - Database logging should be enabled
  - IAM authentication should be configured for RDS instances
  - IAM authentication should be configured for RDS clusters
  - RDS automatic minor version upgrades should be enabled
  - Amazon Aurora clusters should have backtracking enabled
  - RDS DB clusters should be configured for multiple Availability Zones
  - RDS DB clusters should be configured to copy tags to snapshots
  - RDS DB instances should be configured to copy tags to snapshots
  - RDS instances should be deployed in a VPC
  - An RDS event notifications subscription should be configured for critical cluster events
  - An RDS event notifications subscription should be configured for critical database instance events
  - An RDS event notifications subscription should be configured for critical database parameter group events
  - An RDS event notifications subscription should be configured for critical database security group events
  - RDS databases and clusters should not use a database engine default port
  - Amazon Redshift clusters should prohibit public access
  - Connections to Amazon Redshift clusters should be encrypted in transit
  - Amazon Redshift clusters should have automatic snapshots enabled
  - Amazon Redshift clusters should have audit logging enabled
  - Amazon Redshift should have automatic upgrades to major versions enabled
  - Amazon Redshift clusters should use enhanced VPC routing
  - S3 Block Public Access setting should be enabled
  - S3 buckets should prohibit public read access
  - S3 buckets should prohibit public write access
  - S3 buckets should have server-side encryption enabled
  - S3 buckets should deny non-HTTPS requests
  - Amazon S3 permissions granted to other AWS accounts in bucket policies should be restricted
  - Amazon SageMaker notebook instances should not have direct internet access
  - Secrets Manager secrets should have automatic rotation enabled
  - Secrets Manager secrets configured with automatic rotation should rotate successfully
  - Remove unused Secrets Manager secrets
  - Secrets Manager secrets should be rotated within a specified number of days
  - SNS topics should be encrypted at rest using AWS KMS
  - Logging of delivery status should be enabled for notification messages sent to a topic
  - Amazon SQS queues should be encrypted at rest
  - Amazon EC2 instances should be managed by AWS Systems Manager
  - Amazon EC2 instances managed by Systems Manager should have a patch compliance status of COMPLIANT after a patch installation
  - Amazon EC2 instances managed by Systems Manager should have an association compliance status of COMPLIANT
  - SSM documents should not be public
  - AWS WAF Classic global web ACL logging should be enabled

### Dependent Views

AWS Foundational Security Best Practices depends on the following views:

  - view_aws_apigateway_method_settings<sup>*</sup>
  - view_aws_security_group_ingress_rules<sup>*</sup>

  <sup>*</sup> These views are automatically created or updated by this policy.
## AWS Public Egress

### Requirements
AWS Public Egress requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - aws_ec2_instances
  - aws_ec2_route_tables
  - aws_ec2_security_groups
  - aws_lambda_functions
```

### Queries

AWS Public Egress performs the following checks:
  - Find all ec2 instances that have unrestricted access to the internet with a wide open security group and routing
  - All ec2 instances that have unrestricted access to the internet via a security group
  - Find all lambda functions that have unrestricted access to the internet

### Dependent Views

AWS Public Egress depends on the following views:

  - view_aws_security_group_egress_rules<sup>*</sup>

  <sup>*</sup> This view is automatically created or updated by this policy.
## AWS Publicly Available

### Requirements
AWS Publicly Available requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - aws_apigateway_rest_apis
  - aws_apigatewayv2_apis
  - aws_cloudfront_distributions
  - aws_ec2_instances
  - aws_elbv1_load_balancers
  - aws_elbv2_load_balancers
  - aws_rds_instances
  - aws_redshift_clusters
```

### Queries

AWS Publicly Available performs the following checks:
  - Find all API Gateway instances that are publicly accessible
  - Find all API Gateway V2 instances (HTTP and Webhook) that are publicly accessible
  - Find all CloudFront distributions
  - Find all instances with a public IP address
  - Find all Classic ELBs that are Internet Facing
  - Find all ELB V2s that are Internet Facing
  - Amazon Redshift clusters should prohibit public access
  - RDS DB instances should prohibit public access, determined by the PubliclyAccessible configuration
## AWS Unused Resources

### Requirements
AWS Unused Resources requires the following tables to be synced before the policy is executed:

```yaml
tables:
  - aws_acm_certificates
  - aws_apigateway_api_keys
  - aws_backup_vault_recovery_points
  - aws_backup_vaults
  - aws_cloudfront_distributions
  - aws_cloudwatch_alarms
  - aws_directconnect_connections
  - aws_directconnect_lags
  - aws_dynamodb_tables
  - aws_ec2_ebs_volumes
  - aws_ec2_eips
  - aws_ec2_hosts
  - aws_ec2_images
  - aws_ec2_instances
  - aws_ec2_internet_gateways
  - aws_ec2_network_acls
  - aws_ec2_route_tables
  - aws_ec2_security_groups
  - aws_ec2_transit_gateway_attachments
  - aws_ec2_transit_gateways
  - aws_ecr_repositories
  - aws_ecr_repository_images
  - aws_efs_filesystems
  - aws_elbv2_listeners
  - aws_elbv2_load_balancers
  - aws_elbv2_target_groups
  - aws_lightsail_container_service_deployments
  - aws_lightsail_container_services
  - aws_lightsail_disks
  - aws_lightsail_distributions
  - aws_lightsail_load_balancers
  - aws_lightsail_static_ips
  - aws_route53_hosted_zones
  - aws_sns_subscriptions
  - aws_sns_topics
```

### Queries

AWS Unused Resources performs the following checks:
  - Unused ACM certificate
  - Unused API Gateway API key
  - Vaults with no recovery points
  - Disabled CloudFront distribution
  - Disabled CloudWatch alarm
  - Direct Connect connections in "down" state
  - Direct Connect LAGs with no connections
  - DynamoDB table with no items
  - Detached EBS volume
  - Unused EC2 EIP
  - Unused dedicated host
  - Unused own EC2 image
  - Unused internet gateway
  - Unused network access control list
  - Unused EC2 security group
  - Unused route table
  - Unused transit gateway
  - Unused ECR repository
  - Unused EFS filesystem
  - Unused ELB load balancer
  - Unused ELB target group
  - Unused Lightsail container services
  - Unused Lightsail disks
  - Disabled Lightsail distributions
  - Unused Lightsail load balancers
  - Unused Lightsail static IPs
  - Unused Route 53 hosted zones
  - Unused SNS topic
