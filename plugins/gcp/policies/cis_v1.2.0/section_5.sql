\set framework 'cis_v1.2.0'
\echo "Creating CIS V1.2.0 Section 5 Views"
\i views/buckets_permissions.sql
\echo "Executing CIS V1.2.0 Section 5"
\set check_id "5.1"
\echo "Executiong check 5.1"
\i queries/storage/buckets_publicly_accessible.sql
\set check_id "5.2"
\echo "Executiong check 5.2"
\i queries/storage/buckets_without_uniform_bucket_level_access.sql
