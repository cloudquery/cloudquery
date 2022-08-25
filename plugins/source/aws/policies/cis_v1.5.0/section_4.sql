\echo  "Executing CIS V1.5.0 Section 4"
\echo "Creating view_aws_log_metric_filter_and_alarm"
\ir ../views/log_metric_filter_and_alarm.sql
\set check_id '4.1'
\echo "Executing check 4.1"
\ir ../queries/cloudwatch/alarm_unauthorized_api.sql
\set check_id '4.3'
\echo "Executing check 4.3"
\ir ../queries/cloudwatch/alarm_root_account.sql
\set check_id '4.4'
\echo "Executing check 4.4"
\ir ../queries/cloudwatch/alarm_iam_policy_change.sql
\set check_id '4.5'
\echo "Executing check 4.5"
\ir ../queries/cloudwatch/alarm_cloudtrail_config_changes.sql
\set check_id '4.6'
\echo "Executing check 4.6"
\ir ../queries/cloudwatch/alarm_console_auth_failure.sql
\set check_id '4.7'
\echo "Executing check 4.7"
\ir ../queries/cloudwatch/alarm_delete_customer_cmk.sql
\set check_id '4.8'
\echo "Executing check 4.8"
\ir ../queries/cloudwatch/alarm_s3_bucket_policy_change.sql
\set check_id '4.9'
\echo "Executing check 4.9"
\ir ../queries/cloudwatch/alarm_aws_config_changes.sql
\set check_id '4.10'
\echo "Executing check 4.10"
\ir ../queries/cloudwatch/alarm_security_group_changes.sql
\set check_id '4.11'
\echo "Executing check 4.11"
\ir ../queries/cloudwatch/alarm_nacl_changes.sql
\set check_id '4.12'
\echo "Executing check 4.12"
\ir ../queries/cloudwatch/alarm_network_gateways.sql
\set check_id '4.13'
\echo "Executing check 4.13"
\ir ../queries/cloudwatch/alarm_route_table_changes.sql
\set check_id '4.14'
\echo "Executing check 4.14"
\ir ../queries/cloudwatch/alarm_vpc_changes.sql
\set check_id '4.15'
\echo "Executing check 4.15"
\ir ../queries/cloudwatch/alarm_organization_changes.sql
\set check_id '4.16'
\echo "Executing check 4.16"
-- todo add security hub resources to aws provider

