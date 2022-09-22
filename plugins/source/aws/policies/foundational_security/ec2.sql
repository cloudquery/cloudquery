\echo "Creating view_aws_security_group_ingress_rules"
\ir ../views/security_group_ingress_rules.sql

\set check_id 'EC2.1'
\echo "Executing check EC2.1"
\ir ../queries/ec2/ebs_snapshot_permissions_check.sql

\set check_id 'EC2.2'
\echo "Executing check EC2.2"
\ir ../queries/ec2/default_sg_no_access.sql

\set check_id 'EC2.3'
\echo "Executing check EC2.3"
\ir ../queries/ec2/unencrypted_ebs_volumes.sql

\set check_id 'EC2.4'
\echo "Executing check EC2.4"
\ir ../queries/ec2/stopped_more_than_30_days_ago_instances.sql


\set check_id 'EC2.6'
\echo "Executing check EC2.6"
\ir ../queries/ec2/flow_logs_enabled_in_all_vpcs.sql

\set check_id 'EC2.7'
\echo "Executing check EC2.7"
\ir ../queries/ec2/ebs_encryption_by_default_disabled.sql

\set check_id 'EC2.8'
\echo "Executing check EC2.8"
\ir ../queries/ec2/not_imdsv2_instances.sql

\set check_id 'EC2.9'
\echo "Executing check EC2.9"
\ir ../queries/ec2/instances_with_public_ip.sql

\set check_id 'EC2.10'
\echo "Executing check EC2.10"
\ir ../queries/ec2/vpcs_without_ec2_endpoint.sql

\set check_id 'EC2.15'
\echo "Executing check EC2.15"
\ir ../queries/ec2/subnets_that_assign_public_ips.sql

\set check_id 'EC2.16'
\echo "Executing check EC2.16"
\ir ../queries/ec2/unused_acls.sql

\set check_id 'EC2.17'
\echo "Executing check EC2.17"
\ir ../queries/ec2/instances_with_more_than_2_network_interfaces.sql

\set check_id 'EC2.18'
\echo "Executing check EC2.18"
\ir ../queries/ec2/security_groups_with_access_to_unauthorized_ports.sql

\set check_id 'EC2.19'
\echo "Executing check EC2.19"
\ir ../queries/ec2/security_groups_with_open_critical_ports.sql
