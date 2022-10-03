\set framework 'cis_v1.2.0'
\echo "Creating CIS V1.2.0 Section 5 Views"
\ir ../views/buckets_permissions.sql
\echo "Executing CIS V1.2.0 Section 5"
\set check_id '5.1'
\echo "Executing check 5.1"
\ir ../queries/storage/buckets_publicly_accessible.sql
\set check_id '5.2'
\echo "Executing check 5.2"
\ir ../queries/storage/buckets_without_uniform_bucket_level_access.sql
