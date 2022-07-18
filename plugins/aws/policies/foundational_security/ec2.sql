\echo "Creating view_aws_security_group_ingress_rules"
\i views/security_group_ingress_rules.sql

\set check_id 'EC2.1'
\echo "Executing check EC2.1"
\i queries/ec2/ebs_snapshot_permissions_check.sql

\set check_id 'EC2.2'
\echo "Executing check EC2.2"
\i queries/ec2/default_sg_no_access.sql

\set check_id 'EC2.3'
\echo "Executing check EC2.3"
\i queries/ec2/unencrypted_ebs_volumes.sql

\set check_id 'EC2.4'
\echo "Executing check EC2.4"
\i queries/ec2/stopped_more_thant_30_days_ago_instances.sql


\set check_id 'EC2.6'
\echo "Executing check EC2.6"
\i queries/ec2/flow_logs_enabled_in_all_vpcs.sql

\set check_id 'EC2.7'
\echo "Executing check EC2.7"
\i queries/ec2/ebs_encryption_by_default_disabled.sql

\set check_id 'EC2.8'
\echo "Executing check EC2.8"
\i queries/ec2/not_imdsv2_instances.sql

\set check_id 'EC2.9'
\echo "Executing check EC2.9"
\i queries/ec2/instances_with_public_ip.sql

\set check_id 'EC2.10'
\echo "Executing check EC2.10"
\i queries/ec2/vpcs_without_ec2_endpoint.sql

\set check_id 'EC2.15'
\echo "Executing check EC2.15"
\i queries/ec2/subnets_that_assign_public_ips.sql

\set check_id 'EC2.16'
\echo "Executing check EC2.16"
\i queries/ec2/unused_acls.sql

\set check_id 'EC2.17'
\echo "Executing check EC2.17"
\i queries/ec2/instances_with_more_than_2_network_interfaces.sql

\set check_id 'EC2.18'
\echo "Executing check EC2.18"
\i queries/ec2/security_groups_with_access_to_unauthorized_ports.sql

\set check_id 'EC2.19'
\echo "Executing check EC2.19"
\i queries/ec2/security_groups_with_open_critical_ports.sql
