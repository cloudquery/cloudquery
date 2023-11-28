\set check_id 'IAM.1'
\echo "Executing check IAM.1"
\ir ../queries/iam/policies_with_admin_rights.sql

\set check_id 'IAM.2'
\echo "Executing check IAM.2"
\ir ../queries/iam/policies_attached_to_groups_roles.sql

\set check_id 'IAM.3'
\echo "Executing check IAM.3"
\ir ../queries/iam/iam_access_keys_rotated_more_than_90_days.sql

\set check_id 'IAM.4'
\echo "Executing check IAM.4"
\ir ../queries/iam/root_user_no_access_keys.sql

\set check_id 'IAM.5'
\echo "Executing check IAM.5"
\ir ../queries/iam/mfa_enabled_for_console_access.sql

\set check_id 'IAM.6'
\echo "Executing check IAM.6"
\ir ../queries/iam/hardware_mfa_enabled_for_root.sql

\set check_id 'IAM.7'
\echo "Executing check IAM.7"
\ir ../queries/iam/password_policy_strong.sql

\set check_id 'IAM.8'
\echo "Executing check IAM.8"
\ir ../queries/iam/iam_access_keys_unused_more_than_90_days.sql

\set check_id 'IAM.21'
\echo "Executing check IAM.21"
\ir ../queries/iam/wildcard_access_policies.sql
