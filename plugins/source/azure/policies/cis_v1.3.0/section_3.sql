\echo "Executing CIS V1.3.0 Section 3"
\set check_id '3.1'
\echo "Executing check 3.1"
\ir ../queries/storage/secure_transfer_to_storage_accounts_should_be_enabled.sql
\set check_id '3.5'
\echo "Executing check 3.5"
\ir ../queries/storage/no_public_blob_container.sql
\set check_id '3.6'
\echo "Executing check 3.6"
\ir ../queries/storage/default_network_access_rule_is_deny.sql
\set check_id '3.8'
\echo "Executing check 3.8"
\ir ../queries/storage/soft_delete_is_enabled.sql
\set check_id '3.9'
\echo "Executing check 3.9"
\ir ../queries/storage/encrypt_with_cmk.sql
