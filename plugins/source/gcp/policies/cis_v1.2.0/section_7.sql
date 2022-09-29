\set framework 'cis_v1.2.0'
\echo "Executing CIS V1.2.0 Section 7"
\set check_id '7.1'
\echo "Executing check 7.1"
\ir ../queries/bigquery/datasets_publicly_accessible.sql
\set check_id '7.2'
\echo "Executing check 7.2"
\ir ../queries/bigquery/tables_not_encrypted_with_cmek.sql
\set check_id '7.3'
\echo "Executing check 7.3"
\ir ../queries/bigquery/datasets_without_default_cmek.sql
