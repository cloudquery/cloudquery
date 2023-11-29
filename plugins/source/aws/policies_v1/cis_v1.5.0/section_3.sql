\echo "Executing CIS V1.5.0 Section 3"
\set check_id '3.1'
\echo "Executing check 3.1"
\ir ../queries/cloudtrail/enabled_in_all_regions.sql
\set check_id '3.2'
\echo "Executing check 3.2"
\ir ../queries/cloudtrail/log_file_validation_enabled.sql
\set check_id '3.3'
-- todo
\set check_id '3.4'
\echo "Executing check 3.4"
\ir ../queries/cloudtrail/integrated_with_cloudwatch_logs.sql
\set check_id '3.6'
\echo "Executing check 3.6"
\ir ../queries/cloudtrail/bucket_access_logging.sql
\set check_id '3.7'
\echo "Executing check 3.7"
\ir ../queries/cloudtrail/logs_encrypted.sql
\set check_id '3.8'
\echo "Executing check 3.8"
\ir ../queries/kms/rotation_enabled_for_customer_key.sql
\set check_id '3.9'
\echo "Executing check 3.9"
\ir ../queries/ec2/flow_logs_enabled_in_all_vpcs.sql
\set check_id '3.10'
    -- todo cloud_trail_event_selectors data_resources field is wrong
\set check_id '3.11'
    -- todo cloud_trail_event_selectors data_resources field is wrong
