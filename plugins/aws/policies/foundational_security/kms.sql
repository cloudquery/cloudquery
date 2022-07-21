\set check_id 'KMS.1'
\echo "Executing check KMS.1"
\ir ../queries/kms/customer_policy_blocked_kms_actions.sql

\set check_id 'KMS.2'
\echo "Executing check KMS.2"
\ir ../queries/kms/inline_policy_blocked_kms_actions.sql

\set check_id 'KMS.3'
\echo "Executing check KMS.3"
\ir ../queries/kms/cmk_not_scheduled_for_deletion.sql
