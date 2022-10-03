\set ON_ERROR_STOP on
SET TIME ZONE 'UTC';
\set framework 'unused_resources'
-- neat trick to set execution_time if not already set
-- https://stackoverflow.com/questions/32582600/only-set-variable-in-psql-script-if-not-specified-on-the-command-line
\set execution_time :execution_time
SELECT CASE
           WHEN :'execution_time' = ':execution_time' THEN to_char(now(), 'YYYY-MM-dd HH24:MI:SS.US')
           ELSE :'execution_time'
           END AS "execution_time"  \gset
\ir ../create_aws_policy_results.sql

\set check_id 'acm_certificate_unused'
\echo "Executing check acm_certificate_unused"
\ir ../queries/acm/certificates_unused.sql

\set check_id 'apigateway_api_keys_unused'
\echo "Executing check apigateway_api_keys_unused"
\ir ../queries/apigateway/api_keys_disabled.sql

\set check_id 'aws_backup_vaults_unused'
\echo "Executing check aws_backup_vaults_unused"
\ir ../queries/backup/vaults_unused.sql

\set check_id 'aws_cloudfront_distributions_unused'
\echo "Executing check aws_cloudfront_distributions_unused"
\ir ../queries/cloudfront/distributions_disabled.sql

\set check_id 'aws_cloudwatch_alarms_unused'
\echo "Executing check aws_cloudwatch_alarms_unused"
\ir ../queries/cloudwatch/alarm_actions_disabled.sql

\set check_id 'aws_directconnect_connections_unused'
\echo "Executing check aws_directconnect_connections_unused"
\ir ../queries/directconnect/connections_down.sql

\set check_id 'aws_directconnect_lags_unused'
\echo "Executing check aws_directconnect_lags_unused"
\ir ../queries/directconnect/lags_unused.sql

\set check_id 'aws_dynamodb_tables_unused'
\echo "Executing check aws_dynamodb_tables_unused"
\ir ../queries/dynamodb/tables_unused.sql

\set check_id 'aws_ec2_ebs_volumes_unused'
\echo "Executing check aws_ec2_ebs_volumes_unused"
\ir ../queries/ec2/ebs_volumes_detached.sql

\set check_id 'aws_ec2_eips_unused'
\echo "Executing check aws_ec2_eips_unused"
\ir ../queries/ec2/eips_unused.sql

\set check_id 'aws_ec2_hosts_unused'
\echo "Executing check aws_ec2_hosts_unused"
\ir ../queries/ec2/hosts_unused.sql

\set check_id 'aws_ec2_images_unused'
\echo "Executing check aws_ec2_images_unused"
\ir ../queries/ec2/images_unused.sql

\set check_id 'aws_ec2_internet_gateways_unused'
\echo "Executing check aws_ec2_internet_gateways_unused"
\ir ../queries/ec2/internet_gateways_unused.sql

\set check_id 'aws_ec2_network_acls_unused'
\echo "Executing check aws_ec2_network_acls_unused"
\ir ../queries/ec2/network_acls_unused.sql

\set check_id 'aws_ec2_security_groups_unused'
\echo "Executing check aws_ec2_security_groups_unused"
\ir ../queries/ec2/security_groups_unused.sql

\set check_id 'aws_ec2_route_tables_unused'
\echo "Executing check aws_ec2_route_tables_unused"
\ir ../queries/ec2/route_tables_unused.sql

\set check_id 'aws_ec2_transit_gateways_unused'
\echo "Executing check aws_ec2_transit_gateways_unused"
\ir ../queries/ec2/transit_gateways_unused.sql

\set check_id 'aws_ecr_repositories_unused'
\echo "Executing check aws_ecr_repositories_unused"
\ir ../queries/ecr/repositories_unused.sql

\set check_id 'aws_efs_filesystems_unused'
\echo "Executing check aws_efs_filesystems_unused"
\ir ../queries/efs/filesystems_unused.sql

\set check_id 'aws_elb_load_balancers_unused'
\echo "Executing check aws_elb_load_balancers_unused"
\ir ../queries/elb/load_balancers_unused.sql

\set check_id 'aws_elb_target_groups_unused'
\echo "Executing check aws_elb_target_groups_unused"
\ir ../queries/elb/target_groups_unused.sql

\set check_id 'aws_lightsail_container_services_unused'
\echo "Executing check aws_lightsail_container_services_unused"
\ir ../queries/lightsail/container_services_unused.sql

\set check_id 'aws_lightsail_disks_unused'
\echo "Executing check aws_lightsail_disks_unused"
\ir ../queries/lightsail/disks_unused.sql

\set check_id 'aws_lightsail_distributions_unused'
\echo "Executing check aws_lightsail_distributions_unused"
\ir ../queries/lightsail/distributions_unused.sql

\set check_id 'aws_lightsail_load_balancers_unused'
\echo "Executing check aws_lightsail_load_balancers_unused"
\ir ../queries/lightsail/load_balancers_unused.sql

\set check_id 'aws_lightsail_static_ips_unused'
\echo "Executing check aws_lightsail_static_ips_unused"
\ir ../queries/lightsail/static_ips_unused.sql

\set check_id 'aws_route53_hosted_zones_unused'
\echo "Executing check aws_route53_hosted_zones_unused"
\ir ../queries/route53/hosted_zones_unused.sql

\set check_id 'aws_sns_topics_unused'
\echo "Executing check aws_sns_topics_unused"
\ir ../queries/sns/topics_unused.sql
