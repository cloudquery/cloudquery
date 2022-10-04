\set check_id 'S3.1'
\echo "Executing check S3.1"
\ir ../queries/s3/account_level_public_access_blocks.sql

\set check_id 'S3.2'
\echo "Executing check S3.2"
\ir ../queries/s3/publicly_readable_buckets.sql

\set check_id 'S3.3'
\echo "Executing check S3.3"
\ir ../queries/s3/publicly_writable_buckets.sql

\set check_id 'S3.4'
\echo "Executing check S3.4"
\ir ../queries/s3/s3_server_side_encryption_enabled.sql

\set check_id 'S3.5'
\echo "Executing check S3.5"
\ir ../queries/s3/deny_http_requests.sql

\set check_id 'S3.6'
\echo "Executing check S3.6"
\ir ../queries/s3/restrict_cross_account_actions.sql

\set check_id 'S3.8'
\echo "Executing check S3.8"
\ir ../queries/s3/account_level_public_access_blocks.sql
