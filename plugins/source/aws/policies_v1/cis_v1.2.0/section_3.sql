\echo  "Executing CIS V1.2.0 Section 3"
\echo "Creating view_aws_log_metric_filter_and_alarm"
\ir ../views/log_metric_filter_and_alarm.sql
\set check_id '3.1'
\echo "Executing check 3.1"
\ir ../queries/cloudwatch/alarm_unauthorized_api.sql
\set check_id '3.3'
\echo "Executing check 3.3"
\ir ../queries/cloudwatch/alarm_root_account.sql
\set check_id '3.4'
\echo "Executing check 3.4"
\ir ../queries/cloudwatch/alarm_iam_policy_change.sql
\set check_id '3.5'
\echo "Executing check 3.5"
\ir ../queries/cloudwatch/alarm_cloudtrail_config_changes.sql
\set check_id '3.6'
\echo "Executing check 3.6"
\ir ../queries/cloudwatch/alarm_console_auth_failure.sql
\set check_id '3.7'
\echo "Executing check 3.7"
\ir ../queries/cloudwatch/alarm_delete_customer_cmk.sql
\set check_id '3.8'
\echo "Executing check 3.8"
\ir ../queries/cloudwatch/alarm_s3_bucket_policy_change.sql
\set check_id '3.9'
\echo "Executing check 3.9"
\ir ../queries/cloudwatch/alarm_aws_config_changes.sql
\set check_id '3.10'
\echo "Executing check 3.10"
\ir ../queries/cloudwatch/alarm_security_group_changes.sql
\set check_id '3.11'
\echo "Executing check 3.11"
\ir ../queries/cloudwatch/alarm_nacl_changes.sql
\set check_id '3.12'
\echo "Executing check 3.12"
\ir ../queries/cloudwatch/alarm_network_gateways.sql
\set check_id '3.13'
\echo "Executing check 3.13"
\ir ../queries/cloudwatch/alarm_route_table_changes.sql
\set check_id '3.14'
\echo "Executing check 3.14"
\ir ../queries/cloudwatch/alarm_vpc_changes.sql
