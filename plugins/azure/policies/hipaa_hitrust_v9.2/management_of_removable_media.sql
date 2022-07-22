\echo "Management of Removable Media"
\set check_id '0301.09o1Organizational.123 - 09.o - 1'
\echo :check_id
\ir ../queries/sql/data_encryption_off.sql
\set check_id '0302.09o2Organizational.1 - 09.o - 1 (Manual)'
\echo :check_id
\set check_id '0304.09o3Organizational.1 - 09.o - 1'
\echo :check_id
\ir ../queries/datalake/not_encrypted_storage_accounts.sql
\set check_id '0304.09o3Organizational.1 - 09.o - 2'
\echo :check_id
\ir ../queries/sql/managed_instances_without_cmk_at_rest.sql
\set check_id '0304.09o3Organizational.1 - 09.o - 3'
\echo :check_id
\ir ../queries/sql/sqlserver_tde_not_encrypted_with_cmek.sql
