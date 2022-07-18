\set check_id 'CloudTrail.1'
\echo "Executing check CloudTrail.1"
\i queries/cloudtrail/enabled_in_all_regions.sql

\set check_id 'CloudTrail.2'
\echo "Executing check CloudTrail.2"
\i queries/cloudtrail/logs_encrypted.sql

\set check_id 'CloudTrail.4'
\echo "Executing check CloudTrail.4"
\i queries/cloudtrail/log_file_validation_enabled.sql

\set check_id 'CloudTrail.5'
\echo "Executing check CloudTrail.5"
\i queries/cloudtrail/integrated_with_cloudwatch_logs.sql
