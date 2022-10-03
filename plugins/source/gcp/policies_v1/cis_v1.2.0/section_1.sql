\set framework 'cis_v1.2.0'
\echo "Creating CIS V1.2.0 Section 1 Views"
\ir ../views/project_policy_members.sql
\echo "Executing CIS V1.2.0 Section 1"
\set check_id '1.1'
\echo "Executing check 1.1"
\echo "Ensure that corporate login credentials are used (Automated)"
\echo "needs to list folders and organizations which is currently not supported" -- TODO
-- MANUAL
\set check_id '1.2'
\echo "Executing check 1.2"
\echo "Ensure that multi-factor authentication is enabled for all non-service accounts (Manual)"
\ir ../queries/manual.sql
-- MANUAL
\set check_id '1.3'
\echo "Executing check 1.3"
\echo "Ensure that Security Key Enforcement is enabled for all admin accounts (Manual)"
\ir ../queries/manual.sql
\set check_id '1.4'
\echo "Executing check 1.4"
\ir ../queries/iam/managed_service_account_keys.sql
\set check_id '1.5'
\echo "Executing check 1.5"
\ir ../queries/iam/service_account_admin_priv.sql
\set check_id '1.6'
\echo "Executing check 1.6"
\ir ../queries/iam/users_with_service_account_token_creator_role.sql
\set check_id '1.7'
\echo "Executing check 1.7"
\ir ../queries/iam/service_account_keys_not_rotated.sql
\set check_id '1.8'
\echo "Executing check 1.8"
\ir ../queries/iam/separation_of_duties.sql
\set check_id '1.9'
\echo "Executing check 1.9"
\ir ../queries/kms/publicly_accessible.sql
\set check_id '1.10'
\echo "Executing check 1.10"
\ir ../queries/kms/keys_not_rotated_within_90_days.sql
\set check_id '1.11'
\echo "Executing check 1.11"
\ir ../queries/iam/kms_separation_of_duties.sql
-- MANUAL
\set check_id '1.12'
\echo "Executing check 1.12"
\echo "Ensure API keys are not created for a project (Manual)"
\ir ../queries/manual.sql
-- MANUAL
\set check_id '1.13'
\echo "Executing check 1.13"
\echo "Ensure API keys are restricted to use by only specified Hosts and Apps (Manual)"
\ir ../queries/manual.sql
-- MANUAL
\set check_id '1.14'
\echo "Executing check 1.14"
\echo "Ensure API keys are restricted to only APIs that application needs access (Manual)"
\ir ../queries/manual.sql
-- MANUAL
\set check_id '1.15'
\echo "Executing check 1.15"
\echo "Ensure API keys are restricted to only APIs that application needs access (Manual)"
\ir ../queries/manual.sql
