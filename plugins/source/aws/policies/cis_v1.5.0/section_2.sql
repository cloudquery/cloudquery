\echo  "Executing CIS V1.5.0 Section 2"
\set check_id '2.1.1'
\echo "Executing check 2.1.1"
\ir ../queries/s3/s3_server_side_encryption_enabled.sql
\set check_id '2.1.2'
\echo "Executing check 2.1.2"
\ir ../queries/s3/deny_http_requests.sql
\set check_id '2.1.3'
\echo "Executing check 2.1.3"
\ir ../queries/s3/mfa_delete.sql
\set check_id '2.1.4'
    -- manual
\set check_id '2.1.5'
\echo "Executing check 2.1.5"
\ir ../queries/s3/bucket_level_public_access_blocks.sql
\set check_id '2.2.1'
\echo "Executing check 2.2.1"
\ir ../queries/ec2/ebs_encryption_by_default_disabled.sql
\set check_id '2.3.1'
\echo "Executing check 2.3.1"
\ir ../queries/rds/rds_db_instances_should_have_encryption_at_rest_enabled.sql
\set check_id '2.3.2'
\echo "Executing check 2.3.2"
\ir ../queries/rds/rds_automatic_minor_version_upgrades_should_be_enabled.sql
\set check_id '2.3.3'
\echo "Executing check 2.3.3"
\ir ../queries/rds/rds_db_instances_should_prohibit_public_access.sql
\set check_id '2.4.1'
    -- manual