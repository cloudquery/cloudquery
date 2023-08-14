\echo "Executing CIS V1.5.0 Section 3"
\set check_id '3.1'
\echo "Executing check 3.1"
\ir ../queries/cloudtrail/enabled_in_all_regions.sql
\set check_id '3.2'
\echo "Executing check 3.2"
\ir ../queries/cloudtrail/log_file_validation_enabled.sql
\set check_id '3.3'
\echo "Executing check 3.3"
\ir ../queries/cloudtrail/bucket_is_not_public.sql
\set check_id '3.4'
\echo "Executing check 3.4"
\ir ../queries/cloudtrail/integrated_with_cloudwatch_logs.sql
\set check_id '3.5'
\echo "Executing check 3.5"
\ir ../queries/config/enabled_all_regions.sql
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
\echo "Executing check 3.10"
\ir ../queries/cloudtrail/logging_s3_object_writing_event.sql
\set check_id '3.11'
\echo "Executing check 3.11"
\ir ../queries/cloudtrail/logging_s3_object_reading_event.sql
