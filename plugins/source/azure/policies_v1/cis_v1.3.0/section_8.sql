\echo "Executing CIS V1.3.0 Section 8"
\set check_id '8.1'
\echo "Executing check 8.1"
\ir ../queries/keyvault/keys_without_expiration_date.sql
\set check_id '8.2'
\echo "Executing check 8.2"
\ir ../queries/keyvault/secrets_without_expiration_date.sql
\set check_id '8.3'
\echo "Executing check 8.3"
\echo "Check must be done manually"
\set check_id '8.4'
\echo "Executing check 8.4"
\ir ../queries/keyvault/not_recoverable.sql
\set check_id '8.5'
\echo "Executing check 8.5"
\echo "Check must be done manually"
