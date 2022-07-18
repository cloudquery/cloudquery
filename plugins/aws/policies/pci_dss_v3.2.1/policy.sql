\set framework 'pci_dss_v3.2.1'
\set execution_time ''''`date '+%Y-%m-%d %H:%M:%S'`''''::timestamp
\i create_aws_policy_results.sql

\set check_id 'autoscaling.1'
\echo "Executing check autoscaling.1"
\i queries/autoscaling/autoscaling_groups_elb_check.sql

\set check_id 'cloudtrail.1'
\echo "Executing check cloudtrail.1"
\i queries/cloudtrail/logs_encrypted.sql

\set check_id 'cloudtrail.2'
\echo "Executing check cloudtrail.2"
\i queries/cloudtrail/enabled_in_all_regions.sql

\set check_id 'cloudtrail.3'
\echo "Executing check cloudtrail.3"
\i queries/cloudtrail/log_file_validation_enabled.sql

\set check_id 'cloudtrail.4'
\echo "Executing check cloudtrail.4"
\i queries/cloudtrail/integrated_with_cloudwatch_logs.sql

\set check_id 'codebuild.1'
\echo "Executing check codebuild.1"
\i queries/codebuild/check_oauth_usage_for_sources.sql

\set check_id 'codebuild.2'
\echo "Executing check codebuild.2"
\i queries/codebuild/check_environment_variables.sql

\set check_id 'config.1'
\echo "Executing check config.1"
\i queries/config/enabled_all_regions.sql

\echo "Creating view_aws_log_metric_filter_and_alarm"
\i views/log_metric_filter_and_alarm.sql

\set check_id 'cloudwatch.1'
\echo "Executing check cloudwatch.1"
\i queries/cloudwatch/alarm_root_account.sql

\set check_id 'dms.1'
\echo "Executing check dms.1"
\i queries/dms/replication_not_public.sql

\echo "Creating view_aws_security_group_ingress_rules"
\i views/security_group_ingress_rules.sql

\set check_id 'ec2.1'
\echo "Executing check ec2.1"
\i queries/ec2/ebs_snapshot_permissions_check.sql

\set check_id 'ec2.2'
\echo "Executing check ec2.2"
\i queries/ec2/default_sg_no_access.sql


-- This control is retired.
-- Unused EC2 security groups should be removed (Retired)
-- \set check_id 'ec2.3'

\set check_id 'ec2.4'
\echo "Executing check ec2.4"
\i queries/ec2/get_unused_public_ips.sql

\set check_id 'ec2.5'
\echo "Executing check ec2.5"
\i queries/ec2/no_broad_public_ingress_on_port_22.sql

\set check_id 'ec2.6'
\echo "Executing check ec2.6"
\i queries/ec2/flow_logs_enabled_in_all_vpcs.sql

\set check_id 'elbv2.1'
\echo "Executing check elbv2.1"
\i queries/elb/elbv2_redirect_http_to_https.sql

\set check_id 'elasticsearch.1'
\echo "Executing check elasticsearch.1"
\i queries/elasticsearch/elasticsearch_domains_should_be_in_vpc.sql

\set check_id 'elasticsearch.2'
\echo "Executing check elasticsearch.2"
\i queries/elasticsearch/elasticsearch_domains_should_have_encryption_at_rest_enabled.sql

\set check_id 'guardduty enabled in all enabled regions'
\echo "Executing check guardduty enabled in all enabled regions"
\i queries/guardduty/detector_enabled.sql

\set check_id 'iam.1'
\echo "Executing check iam.1"
\i queries/iam/root_user_no_access_keys.sql

\set check_id 'iam.2'
\echo "Executing check iam.2"
\i queries/iam/policies_attached_to_groups_roles.sql

\set check_id 'iam.3'
\echo "Executing check iam.3"
\i queries/iam/no_star.sql

\set check_id 'iam.4'
\echo "Executing check iam.4"
\i queries/iam/hardware_mfa_enabled_for_root.sql

\set check_id 'iam.5'
\echo "Executing check iam.5"
\i queries/iam/mfa_enabled_for_root.sql

\set check_id 'iam.6'
\echo "Executing check iam.6"
\i queries/iam/mfa_enabled_for_console_access.sql

\set check_id 'iam.7'
\echo "Executing check iam.7"
\i queries/iam/unused_creds_disabled.sql

\set check_id 'iam.8'
\echo "Executing check iam.8"
\i queries/iam/password_policy_strong.sql

\set check_id 'kms.1'
\echo "Executing check kms.1"
\i queries/kms/rotation_enabled_for_customer_key.sql

\set check_id 'lambda.1'
\echo "Executing check lambda.1"
\i queries/lambda/lambda_function_prohibit_public_access.sql

\set check_id 'lambda.2'
\echo "Executing check lambda.2"
\i queries/lambda/lambda_function_in_vpc.sql

\set check_id 'rds.1'
\echo "Executing check rds.1"
\i queries/rds/snapshots_should_prohibit_public_access.sql

\set check_id 'rds.2'
\echo "Executing check rds.2"
\i queries/rds/rds_db_instances_should_prohibit_public_access.sql

\set check_id 'redshift.1'
\echo "Executing check redshift.1"
\i queries/redshift/cluster_publicly_accessible.sql

\set check_id 's3.1'
\echo "Executing check s3.1"
\i queries/s3/publicly_writable_buckets.sql

\set check_id 's3.2'
\echo "Executing check s3.2"
\i queries/s3/publicly_readable_buckets.sql

\set check_id 's3.3'
\echo "Executing check s3.3"
\i queries/s3/s3_cross_region_replication.sql

\set check_id 's3.4'
\echo "Executing check s3.4"
\i queries/s3/s3_server_side_encryption_enabled.sql

\set check_id 's3.5'
\echo "Executing check s3.5"
\i queries/s3/deny_http_requests.sql

\set check_id 's3.6'
\echo "Executing check s3.6"
\i queries/s3/account_level_public_access_blocks.sql

\set check_id 'sagemaker.1'
\echo "Executing check sagemaker.1"
\i queries/sagemaker/sagemaker_notebook_instance_direct_internet_access_disabled.sql

\set check_id 'secretmanager.1'
\echo "Executing check secretmanager.1"
\i queries/secretsmanager/secrets_should_have_automatic_rotation_enabled.sql

\set check_id 'secretmanager.2'
\echo "Executing check secretmanager.2"
\i queries/secretsmanager/secrets_configured_with_automatic_rotation_should_rotate_successfully.sql

\set check_id 'secretmanager.3'
\echo "Executing check secretmanager.3"
\i queries/secretsmanager/remove_unused_secrets_manager_secrets.sql

\set check_id 'secretmanager.4'
\echo "Executing check secretmanager.4"
\i queries/secretsmanager/secrets_should_be_rotated_within_a_specified_number_of_days.sql

\set check_id 'ssm.1'
\echo "Executing check ssm.1"
\i queries/ssm/instances_should_have_patch_compliance_status_of_compliant.sql

\set check_id 'ssm.2'
\echo "Executing check ssm.2"
\i queries/ssm/instances_should_have_association_compliance_status_of_compliant.sql

\set check_id 'ssm.3'
\echo "Executing check ssm.3"
\i queries/ssm/ec2_instances_should_be_managed_by_ssm.sql

\set check_id 'waf.1'
\echo "Executing check waf.1"
\i queries/wafv2/wafv2_web_acl_logging_should_be_enabled.sql
